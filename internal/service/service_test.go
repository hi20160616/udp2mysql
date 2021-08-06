package service

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/hi20160616/udp2mysql/api/udp2mysql/v1"
)

var us = InitUDPService()

func TestListUDPPackets(t *testing.T) {
	udps, err := us.ListUDPPackets(context.Background(), &v1.ListUDPPacketsRequest{})
	if err != nil {
		t.Error(err)
	}
	for _, e := range udps.UdpPackets {
		fmt.Println(e)
	}
}

func TestGetUDPPacket(t *testing.T) {
	udp, err := us.GetUDPPacket(context.Background(),
		&v1.GetUDPPacketRequest{Name: "ee4056c69f21b924a470c8a56a4e0f5b"})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(udp)
}
