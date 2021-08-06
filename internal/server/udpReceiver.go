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
func NewUDPReceiver() (*UDPReceiver, error) {
	s, err := net.ResolveUDPAddr("udp4", configs.V.UDPSender.Addr)
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
		buf:     make([]byte, configs.V.UDPSender.BufSize),
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
			if err = ur.deal(n, addr); err != nil {
				log.Println(err)
			}
		}
	}
}

// deal with udp packets
func (ur *UDPReceiver) deal(n int, addr *net.UDPAddr) error {
	// get bytes receive
	u := ur.buf[0:n]
	// TODO: dail to ms
	// TODO: send udp packets to ms

	// just for test workflow
	fmt.Print("-> ", string(u), "\n")
	reply := []byte(time.Now().String())
	fmt.Printf("Server reply data: %s\n", reply)
	_, err := ur.conn.WriteToUDP(reply, addr)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UDPReceiver) Stop(ctx context.Context) error {
	return ur.conn.Close()
}
