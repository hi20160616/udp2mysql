package biz

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type UDPPacket struct {
	Name, Id, Title, Content string
	UpdateTime               *timestamppb.Timestamp
}

type UDPPacketRepo interface {
	ListUDPPackets(ctx context.Context) ([]*UDPPacket, error)
	// name means the last part of url path like `/udp/uuid`, uuid is name
	GetUDPPackets(ctx context.Context, name string) (*UDPPacket, error)
	CreateUDPPacket(ctx context.Context, udp *UDPPacket) (*UDPPacket, error)
	UpdateUDPPacket(ctx context.Context, udp *UDPPacket) (*UDPPacket, error)
	// name means the last part of url path like `/udp/uuid`, uuid is name
	DeleteUDPPacket(ctx context.Context, name string) error
}

type UDPPacketUsecase struct {
	repo UDPPacketRepo
}

func NewUDPPacketUsecase(repo UDPPacketRepo) *UDPPacketUsecase {
	return &UDPPacketUsecase{repo: repo}
}
