package mariadb

import (
	"database/sql"

	"github.com/hi20160616/udp2mysql/configs"
)

type Client struct {
	db        *sql.DB
	UDPPacket *UDPPacketClient
	Err       error
}

type UDPPacketClient struct {
	db *sql.DB
}

func open() (*sql.DB, error) {
	return sql.Open(configs.V.Database.Driver, configs.V.Database.Source)
}

func NewClient() *Client {
	db, err := open()
	return &Client{db, &UDPPacketClient{db}, err}
}
