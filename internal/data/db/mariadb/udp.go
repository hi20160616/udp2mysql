package mariadb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UDPPacket struct {
	ID, Name, Title, Content string
	UpdateTime               timestamppb.Timestamp
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
	up := &UDPPacket{}
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
