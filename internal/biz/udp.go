package biz

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type UDPPacket struct {
	Name, Id, Title, Content string
	UpdateTime               *timestamppb.Timestamp
}

type UDPPackets struct {
	UdpPackets    []*UDPPacket
	NextPageToken string
}

type UDPPacketRepo interface {
	ListUDPPackets(ctx context.Context) (*UDPPackets, error)
	// name means the last part of url path like `/udp/uuid`, uuid is name
	GetUDPPacket(ctx context.Context, name string) (*UDPPacket, error)
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

func (uc *UDPPacketUsecase) List(ctx context.Context) (*UDPPackets, error) {
	return uc.repo.ListUDPPackets(ctx)
}

func (uc *UDPPacketUsecase) Get(ctx context.Context, name string) (*UDPPacket, error) {
	return uc.repo.GetUDPPacket(ctx, name)
}

func (uc *UDPPacketUsecase) Create(ctx context.Context, udp *UDPPacket) (*UDPPacket, error) {
	return uc.repo.CreateUDPPacket(ctx, udp)
}

func (uc *UDPPacketUsecase) Update(ctx context.Context, udp *UDPPacket) (*UDPPacket, error) {
	return uc.repo.UpdateUDPPacket(ctx, udp)
}

func (uc *UDPPacketUsecase) Delete(ctx context.Context, name string) error {
	return uc.repo.DeleteUDPPacket(ctx, name)
}
