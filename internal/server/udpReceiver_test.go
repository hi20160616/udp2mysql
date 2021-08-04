package server

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/hi20160616/udp2mysql/configs"
)

func TestUDPReceiver(t *testing.T) {
	ur, err := NewUDPReceiver(configs.V.RemoteAddr, 1024)
	if err != nil {
		t.Error(err)
	}
	go func() {
		if err := ur.Start(context.Background()); err != nil {
			t.Error(err)
		}
	}()

	c, err := net.DialUDP("udp4", nil, ur.udpAddr)
	if err != nil {
		t.Error(err)
	}
	defer c.Close()
	data := []byte("This is test text!")
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

	if err := ur.Stop(context.Background()); err != nil {
		t.Error(err)
	}
}
