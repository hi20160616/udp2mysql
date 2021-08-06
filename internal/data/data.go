package data

import (
	"github.com/hi20160616/udp2mysql/internal/biz"
	"github.com/hi20160616/udp2mysql/internal/data/db/mariadb"
)

var _ biz.UDPPacketRepo = new(udpPacketRepo)

type Data struct {
	DBClient *mariadb.Client
}
