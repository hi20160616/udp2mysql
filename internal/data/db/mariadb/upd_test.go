package mariadb

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func TestUpdateUDPPacket(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	upkt := &UDPPacket{
		ID:      "2c8906da0775c93fdcb57b401e622e18",
		Name:    "test7 name",
		Title:   "test7 title",
		Content: "test7 content",
	}
	if err := c.UDPPacket.Update(context.Background(), upkt); err != nil {
		t.Error(err)
	}
	ps := [4]string{"id", "=", "2c8906da0775c93fdcb57b401e622e18"}
	got, err := c.UDPPacket.Query().Where(ps).First(context.Background())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(got)
}

func TestDeleteUDPPacket(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	id := "2c8906da0775c93fdcb57b401e622e18"
	if err := c.UDPPacket.Delete(context.Background(), id); err != nil {
		t.Error(err)
	}
	got, err := c.UDPPacket.Query().
		Where([4]string{"id", "=", id}).
		First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
	if got != nil {
		t.Error(errors.New("Delete failed."))
	}
}
