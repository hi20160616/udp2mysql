package mariadb

import "google.golang.org/protobuf/types/known/timestamppb"

type UDPPacket struct {
	ID, Title, Content string
	UpdateTime         timestamppb.Timestamp
}

type UDPPacketQuery struct {
	limit    int
	offset   int64
	query    string
	keywords []string
}

func (uq *UDPPacketQuery) Query() ([]*UDPPacket, error) {
	return nil, nil
}

func (uq *UDPPacketQuery) List() ([]*UDPPacket, error) {
	return nil, nil
}