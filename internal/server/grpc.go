package server

import (
	"net"
	"time"

	pb "github.com/hi20160616/udp2mysql/api/udp2mysql/v1"
	"github.com/hi20160616/udp2mysql/configs"
	"github.com/hi20160616/udp2mysql/internal/service"
	"google.golang.org/grpc"
)

type GRPC struct {
	s *grpc.Server
	l *net.Listener
}

func NewGRPCServer(udp *service.UDPService) (*GRPC, error) {
	t, err := time.ParseDuration(configs.V.API.GRPC.Timeout)
	if err != nil {
		return nil, err
	}
	opts := []grpc.ServerOption{
		grpc.MaxMsgSize(1 << 30),
		grpc.ConnectionTimeout(t),
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterUDPPacketApiServer(srv, udp)
	l, err := net.Listen("tcp", configs.V.API.GRPC.Addr)
	if err != nil {
		return nil, err
	}
	return &GRPC{srv, &l}, nil
}
