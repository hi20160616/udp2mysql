package mariadb

import (
	"database/sql"
	"fmt"

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

func (uc *UDPPacketClient) Query() *UDPPacketQuery {
	return &UDPPacketQuery{}
}

func open() (*sql.DB, error) {
	fmt.Println(configs.V.Database.Driver)
	return sql.Open(configs.V.Database.Driver, configs.V.Database.Source)
}

func NewClient() *Client {
	db, err := open()
	return &Client{db, &UDPPacketClient{db}, err}
}
