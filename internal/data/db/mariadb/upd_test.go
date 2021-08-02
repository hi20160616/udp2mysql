package mariadb

import (
	"context"
	"crypto/md5"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	upkt1 := &UDPPacket{
		ID:         fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(time.Now().Nanosecond())))),
		Name:       "test1 name",
		Title:      "test1 title",
		Content:    "test1 content",
		UpdateTime: time.Now(),
	}
	upkt2 := &UDPPacket{
		ID:         fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(time.Now().Nanosecond())))),
		Name:       "test2 name",
		Title:      "test2 title",
		Content:    "test2 content",
		UpdateTime: time.Now(),
	}
	upkt3 := &UDPPacket{
		ID:         fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(time.Now().Nanosecond())))),
		Name:       "test3 name",
		Title:      "test3 title",
		Content:    "test3 content",
		UpdateTime: time.Now(),
	}
	if err := c.UDPPacket.Insert(context.Background(), upkt1); err != nil {
		t.Error(err)
	}
	if err := c.UDPPacket.Insert(context.Background(), upkt2); err != nil {
		t.Error(err)
	}
	if err := c.UDPPacket.Insert(context.Background(), upkt3); err != nil {
		t.Error(err)
	}
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
		{"name", "like", "test"},
		// {"name", "like", "test", "and"},
		// {"title", "like", "title", "and"},
		// {"content", "=", "test content"},
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
