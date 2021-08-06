package main

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	pb "github.com/hi20160616/udp2mysql/api/udp2mysql/v1"
	"github.com/hi20160616/udp2mysql/configs"
	"google.golang.org/grpc"
)

func TestUDPReceiver(t *testing.T) {
	s, err := net.ResolveUDPAddr("udp4", configs.V.UDPSender.Addr)
	if err != nil {
		t.Error(err)
	}
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		data := []byte(fmt.Sprintf("This is test text %d!", i+1))
		_, err = c.Write(data)
		if err != nil {
			t.Error(err)
		}
		buf := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buf)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("Client receive: %s\n", string(buf[:n]))
		time.Sleep(time.Second)
	}
}

func TestGRPCServer(t *testing.T) {
	tt, err := time.ParseDuration(configs.V.API.GRPC.Timeout)
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), tt)
	defer cancel()

	// Set up a connection to the server
	conn, err := grpc.Dial(configs.V.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	c := pb.NewUDPPacketApiClient(conn)

	// Contact the server and print out its response.
	id := "3a181362ab507cc674b94265a48c0235"
	u, err := c.GetUDPPacket(ctx, &pb.GetUDPPacketRequest{Name: id})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
	udps, err := c.ListUDPPackets(ctx, &pb.ListUDPPacketsRequest{})
	if err != nil {
		t.Error(err)
	}
	for _, e := range udps.UdpPackets {
		fmt.Println(e)
	}
}
