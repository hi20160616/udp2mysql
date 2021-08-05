package main

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/hi20160616/udp2mysql/configs"
)

func TestUDPReceiver(t *testing.T) {
	s, err := net.ResolveUDPAddr("udp4", configs.V.RemoteAddr)
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
