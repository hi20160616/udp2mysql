package data

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hi20160616/udp2mysql/internal/biz"
	"github.com/hi20160616/udp2mysql/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateUDPPacket(t *testing.T) {
	tc := &biz.UDPPacket{
		Id:         fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(time.Now().Nanosecond())))),
		Name:       "test4 name",
		Title:      "test4 title",
		Content:    "test4 content",
		UpdateTime: timestamppb.Now(),
	}
	data := mariadb.NewClient()
	ur := NewUDPPacketRepo(&Data{dbClient: data})
	up, err := ur.CreateUDPPacket(context.Background(), tc)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(up)
}

func TestListUDPPackets(t *testing.T) {
	data := mariadb.NewClient()
	ur := NewUDPPacketRepo(&Data{dbClient: data})
	us, err := ur.ListUDPPackets(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, e := range us.UdpPackets {
		fmt.Println(e)
	}
}

func TestGetUDPPacket(t *testing.T) {
	data := mariadb.NewClient()
	ur := NewUDPPacketRepo(&Data{dbClient: data})
	u, err := ur.GetUDPPacket(context.Background(), "test4 name")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
}

func TestDeleteUDPPacket(t *testing.T) {
	data := mariadb.NewClient()
	ur := NewUDPPacketRepo(&Data{dbClient: data})
	id := "28a228e245adc300a7439f3f7b1a2969"
	err := ur.DeleteUDPPacket(context.Background(), id)
	if err != nil {
		t.Error(err)
	}
	u, err := ur.GetUDPPacket(context.Background(), "test6 name")
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
	if u != nil {
		fmt.Println(u)
		t.Error(errors.New("Delete failed"))
	}
}
