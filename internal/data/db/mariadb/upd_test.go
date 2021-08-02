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

func TestWhere(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	ps := [][4]string{
		{"name", "like", "test", "and"},
		{"title", "like", "title", "and"},
		{"content", "=", "test content"},
	}
	got, err := c.UDPPacket.Query().Where(ps...).All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, v := range got.Collection {
		fmt.Println(v)
	}

}

func TestPrepareQuery(t *testing.T) {
	uq := &UDPPacketQuery{}
	uq.query = "SELECT * FROM udp_packets"
	uq.Where([4]string{"name", "like", "test"})
	if err := uq.prepareQuery(context.Background()); err != nil {
		t.Error(err)
	}
	fmt.Println(uq)
}
