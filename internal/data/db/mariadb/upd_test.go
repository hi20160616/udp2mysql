package mariadb

import (
	"context"
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	got, err := c.UDPPacket.Create().Save(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	fmt.Println(got)
}

func TestListUDPPackets(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	got, err := c.UDPPacket.Query().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, v := range got.Collection {
		fmt.Println(v)
	}
}
