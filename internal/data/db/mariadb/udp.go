package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
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
	limit      *int
	offset     *int
	query      string
	predicates [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order      string
	args       []interface{}
	keywords   []string
}

func (uc *UDPPacketClient) Insert(ctx context.Context, upkt *UDPPacket) error {
	q := "INSERT INTO udp_packets(id, name, title, content, update_time) VALUES(?,?,?,?,?)" +
		" ON DUPLICATE KEY UPDATE id=?, name=?, title=?, content=?, update_time=?"
	uq := &UDPPacketQuery{db: uc.db, query: q}
	_, err := uq.db.Exec(uq.query,
		upkt.ID, upkt.Name, upkt.Title, upkt.Content, upkt.UpdateTime,
		upkt.ID, upkt.Name, upkt.Title, upkt.Content, upkt.UpdateTime)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Save error")
	}
	return nil
}

func (uc *UDPPacketClient) Query() *UDPPacketQuery {
	return &UDPPacketQuery{
		db:    uc.db,
		query: "SELECT * FROM udp_packets",
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

func (uq *UDPPacketQuery) First(ctx context.Context) (*UDPPacket, error) {
	nodes, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes.Collection) == 0 {
		return nil, ErrNotFound
	}
	return nodes.Collection[0], nil
}

// ps: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
func (uq *UDPPacketQuery) Where(ps ...[4]string) *UDPPacketQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

func (uq *UDPPacketQuery) Order(condition string) *UDPPacketQuery {
	uq.order = condition
	return uq
}

func (uq *UDPPacketQuery) Limit(limit int) *UDPPacketQuery {
	uq.limit = &limit
	return uq
}

func (uq *UDPPacketQuery) Offset(offset int) *UDPPacketQuery {
	uq.offset = &offset
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
	if uq.order != "" {
		uq.query += " ORDER BY ?"
		uq.args = append(uq.args, uq.order)
	}
	if uq.limit != nil {
		uq.query += " LIMIT ?"
		a := strconv.Itoa(*uq.limit)
		uq.args = append(uq.args, a)
	}
	if uq.offset != nil {
		uq.query += ", ?"
		a := strconv.Itoa(*uq.offset)
		uq.args = append(uq.args, a)
	}
	fmt.Println(uq.query)
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
