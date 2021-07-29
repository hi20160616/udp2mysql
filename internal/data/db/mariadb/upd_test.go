package mariadb

import (
	"context"
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	c := NewClient()
	got := c.UDPPacket.Create()
	fmt.Println(got)
}

func TestSave(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Error(c.Err)
		return
	}
	got, err := c.UDPPacket.Create().Save(context.Background())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(got)
}
