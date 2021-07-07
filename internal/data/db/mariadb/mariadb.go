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
}

func NewUDPPacketClient() *UDPPacketClient {
	return &UDPPacketClient{}
}

func (uc *UDPPacketClient) Query() *UDPPacketQuery {
	return &UDPPacketQuery{}
}

func open() (*sql.DB, error) {
	return sql.Open(configs.V.Data.Driver, configs.V.Data.Source)
}

func NewClient() *Client {
	db, err := open()
	return &Client{db, nil, err}
}

func (c *Client) init() {
	c.UDPPacket = NewUDPPacketClient()
}
