package mariadb

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UDPPacket struct {
	ID, Name, Title, Content string
	UpdateTime               time.Time
}

type UDPPackets struct {
	Collection []*UDPPacket
}

type UDPPacketQuery struct {
	db         *sql.DB
	limit      int
	offset     int64
	query      string
	predicates [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	args       []interface{}
	keywords   []string
}

func (uc *UDPPacketClient) Create() *UDPPacketQuery {
	fmt.Println("Create UDPPacket at mariadb.go")
	q := "INSERT INTO udp_packets(id, name, title, content, update_time) VALUES(?,?,?,?,?)" +
		" ON DUPLICATE KEY UPDATE id=?, name=?, title=?, content=?, update_time=?"
	return &UDPPacketQuery{
		db:    uc.db,
		query: q,
	}
}

func (uq *UDPPacketQuery) Save(ctx context.Context) (*UDPPacket, error) {
	up := &UDPPacket{
		ID:         fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(time.Now().Nanosecond())))),
		Name:       "test name",
		Title:      "test title",
		Content:    "test content",
		UpdateTime: timestamppb.Now().AsTime(),
	}
	_, err := uq.db.Exec(uq.query,
		up.ID, up.Name, up.Title, up.Content, up.UpdateTime,
		up.ID, up.Name, up.Title, up.Content, up.UpdateTime)
	if err != nil {
		return nil, errors.WithMessage(err, "mariadb: Save error")
	}
	return nil, nil
}

func (uc *UDPPacketClient) Query() *UDPPacketQuery {
	q := "SELECT * FROM udp_packets"
	return &UDPPacketQuery{
		db:    uc.db,
		query: q,
	}
}

func (uq *UDPPacketQuery) All(ctx context.Context) (*UDPPackets, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := uq.db.Query(uq.query, uq.args...)
	// rows, err := uq.db.Query("SELECT * FROM udp_packets WHERE name like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkUDPPacket(rows)
}

// ps: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
func (uq *UDPPacketQuery) Where(ps ...[4]string) *UDPPacketQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

func (uq *UDPPacketQuery) prepareQuery(ctx context.Context) error {
	if uq.predicates != nil {
		uq.query += " WHERE "
		for _, p := range uq.predicates {
			uq.query += fmt.Sprintf(" %s %s ? %s", p[0], p[1], p[3])
			if strings.ToLower(p[1]) == "like" {
				p[2] = fmt.Sprintf("%%%s%%", p[2])
			} else {
				p[2] = fmt.Sprintf("%s", p[2])
			}
			uq.args = append(uq.args, p[2])
		}
	}
	return nil
}

func mkUDPPacket(rows *sql.Rows) (*UDPPackets, error) {
	var id, name, title, content sql.NullString
	var update_time sql.NullTime
	var udp_packets = &UDPPackets{}
	for rows.Next() {
		if err := rows.Scan(&id, &name, &title, &content, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkUDPPacket rows.Scan error")
		}
		udp_packets.Collection = append(udp_packets.Collection, &UDPPacket{
			ID:         id.String,
			Name:       name.String,
			Title:      title.String,
			Content:    content.String,
			UpdateTime: update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkUDPPacket error")
	}
	return udp_packets, nil
}
