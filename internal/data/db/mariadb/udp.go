package mariadb

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UDPPacket struct {
	ID, Name, Title, Content string
	UpdateTime               time.Time
}

type UDPPacketQuery struct {
	db       *sql.DB
	limit    int
	offset   int64
	query    string
	keywords []string
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

func (uq *UDPPacketQuery) Query() ([]*UDPPacket, error) {
	return nil, nil
}

func (uq *UDPPacketQuery) List() ([]*UDPPacket, error) {
	return nil, nil
}
