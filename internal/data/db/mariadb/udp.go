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
	query      string
	limit      *int
	offset     *int
	order      string
	fields     []string
	predicates [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	args       []string
	keywords   []string
}

var Columns = []string{"*", "id", "name", "title", "content", "update_time"}
var Table = "udp_packets"

func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
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

func (uc *UDPPacketClient) NewQuery() *UDPPacketQuery {
	return &UDPPacketQuery{
		db: uc.db,
	}
}

func (uq *UDPPacketQuery) Select(field string, fields ...string) *UDPPacketQuery {
	uq.fields = append([]string{field}, fields...)
	return uq
}

// ps: [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
func (uq *UDPPacketQuery) Where(ps ...[4]string) *UDPPacketQuery {
	// for _, p := range ps {
	//         if !ValidColumn(p[0]) {
	//                 return errors.WithMessagef(ValidationError, "Where: invalid field %q for query", p[0])
	//         }
	//         uq.predicates += fmt.Sprintf(" %s%s%s %s ", p[0], p[1], p[2], p[3])
	// }
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

func (uq *UDPPacketQuery) All(ctx context.Context) (*UDPPackets, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := uq.db.Query(uq.query, uq.args)
	if err != nil {
		return nil, err
	}
	return mkUDPPacket(rows)
}

func (uq *UDPPacketQuery) prepareQuery(ctx context.Context) error {
	if uq.fields == nil {
		uq.fields = append(uq.fields, "*")
	}
	for _, f := range uq.fields {
		if !ValidColumn(f) {
			return errors.WithMessagef(ErrValidation, "invalid field %q for query", f)
		}
	}
	if uq.fields != nil {
		uq.query = "SELECT ? FROM " + Table
		fields := strings.Join(uq.fields, ",")
		if strings.Contains(fields, "*") {
			fields = "*"
		}
		uq.args = append(uq.args, fields)
	}
	// predicate pattern: ["name", "=", "jack", "and"]
	if uq.predicates != nil {
		uq.query += " WHERE ?"
		predicates := func() string {
			a := ""
			if uq.predicates == nil {
				return ""
			}
			for _, p := range uq.predicates {
				if strings.ToLower(p[1]) == "like" {
					p[2] = fmt.Sprintf("\"%%%s%%\"", p[2])
				} else {
					p[2] = fmt.Sprintf("\"%s\"", p[2])
				}
				a = strings.Join(p[:3], "") + " " + p[3]
				// if i == len(uq.predicates)-1 { // last element
				//         a += fmt.Sprintf("%s%s%s", p[0], p[1], p[2])
				// } else {
				//         a += fmt.Sprintf("%s%s%s %s", p[0], p[1], p[2], p[3])
				// }
			}
			return a
		}()
		if predicates != "" {
			uq.args = append(uq.args, predicates)
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
