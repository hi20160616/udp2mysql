package server

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/hi20160616/udp2mysql/api/udp2mysql/v1"
	"github.com/hi20160616/udp2mysql/configs"
	"github.com/hi20160616/udp2mysql/internal/service"
	"google.golang.org/grpc"
)

type GRPC struct {
	s *grpc.Server
	l net.Listener
}

func NewGRPCServer() (*GRPC, error) {
	t, err := time.ParseDuration(configs.V.API.GRPC.Timeout)
	if err != nil {
		return nil, err
	}
	opts := []grpc.ServerOption{
		// grpc.MaxMsgSize(1 << 30),
		grpc.ConnectionTimeout(t),
	}
	s := grpc.NewServer(opts...)
	us := service.InitUDPService()
	pb.RegisterUDPPacketApiServer(s, us)
	l, err := net.Listen("tcp", configs.V.API.GRPC.Addr)
	if err != nil {
		return nil, err
	}
	return &GRPC{s, l}, nil
}

func (gs *GRPC) Start(ctx context.Context) error {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e)
			PanicLog(e)
		}
	}()
	defer gs.l.Close()
	return gs.s.Serve(gs.l)
}

func (gs *GRPC) Stop(ctx context.Context) error {
	gs.s.GracefulStop()
	return nil
}
