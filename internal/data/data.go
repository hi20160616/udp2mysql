package data

import (
	"github.com/hi20160616/udp2mysql/internal/data/db/mariadb"
)

type Data struct {
	dbClient *mariadb.Client
}
