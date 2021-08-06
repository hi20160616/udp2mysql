package server

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestUDPReceiver(t *testing.T) {
	ur, err := NewUDPReceiver()
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
	if err := ur.Stop(context.Background()); err != nil {
		t.Error(err)
	}
}
