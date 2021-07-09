package mariadb

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	c := NewClient()
	got := c.UDPPacket.Create()
	fmt.Println(got)
}
