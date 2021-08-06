package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hi20160616/udp2mysql/configs"
	myerr "github.com/hi20160616/udp2mysql/errors"
)

type UDPReceiver struct {
	conn    *net.UDPConn
	udpAddr *net.UDPAddr
	buf     []byte
}

// addr seems like "127.0.0.1:1234"
func NewUDPReceiver(addr string, bufSize int) (*UDPReceiver, error) {
	s, err := net.ResolveUDPAddr("udp4", configs.V.RemoteAddr)
	if err != nil {
		return nil, err
	}
	l, err := net.ListenUDP("udp4", s)
	if err != nil {
		return nil, err
	}
	return &UDPReceiver{
		conn:    l,
		udpAddr: s,
		buf:     make([]byte, bufSize),
	}, nil
}

func (ur *UDPReceiver) Start(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e)
			myerr.PanicLog(e)
		}
	}()
	defer ur.conn.Close()

	for {
		n, addr, err := ur.conn.ReadFromUDP(ur.buf)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(30 * time.Second):
			return fmt.Errorf("context was not done immediately")
		default:
			fmt.Print("-> ", string(ur.buf[0:n]), "\n")
			reply := []byte(time.Now().String())
			fmt.Printf("Server reply data: %s\n", reply)
			_, err = ur.conn.WriteToUDP(reply, addr)
			if err != nil {
				log.Printf("%v", err)
			}
		}
	}
}

func (ur *UDPReceiver) Stop(ctx context.Context) error {
	return ur.conn.Close()
}
