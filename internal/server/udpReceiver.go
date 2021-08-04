package server

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type UDPReceiver struct {
	udpAddr *net.UDPAddr
	buf     []byte
}

// addr seems like "127.0.0.1:1234"
func NewUDPReceiver(addr string, bufSize int) (*UDPReceiver, error) {
	s, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		return nil, err
	}
	return &UDPReceiver{
		udpAddr: s,
		buf:     make([]byte, bufSize),
	}, nil
}

var done = make(chan struct{}, 1)

func (ur *UDPReceiver) Start(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e)
			PanicLog(e)
		}
	}()

	l, err := net.ListenUDP("udp4", ur.udpAddr)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		n, addr, err := l.ReadFromUDP(ur.buf)
		fmt.Print("-> ", string(ur.buf[0:n-1]))

		reply := []byte(time.Now().String())
		fmt.Printf("\nServer reply data: %s\n", reply)
		_, err = l.WriteToUDP(reply, addr)
		if err != nil {
			log.Printf("%v", err)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-done:
			return nil
		}
	}
}

func (ur *UDPReceiver) Stop(ctx context.Context) error {
	done <- struct{}{}
	return ctx.Err()
}

func PanicLog(_err error) error {
	filePath := "./PanicLog.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("[" + time.Now().Format(time.RFC3339) + "]--------------------------------------\n")
	write.WriteString(_err.Error() + "\n")
	write.Flush()
	return nil
}
