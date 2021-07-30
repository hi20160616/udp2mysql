package data

import (
	"context"

	"github.com/hi20160616/udp2mysql/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.UDPPacketRepo = new(udpPacketRepo)

type udpPacketRepo struct {
	data *Data
}

func NewUDPPacketRepo(data *Data) biz.UDPPacketRepo {
	return &udpPacketRepo{
		data: data,
	}
}

func (ur *udpPacketRepo) ListUDPPackets(ctx context.Context) ([]*biz.UDPPacket, error) {
	us, err := ur.data.db.UDPPacket.NewQuery().All(ctx)
	if err != nil {
		return nil, err
	}
	bus := make([]*biz.UDPPacket, 0)
	for _, u := range us.Collection {
		bus = append(bus, &biz.UDPPacket{
			Id:         u.ID,
			Title:      u.Title,
			Content:    u.Content,
			UpdateTime: timestamppb.New(u.UpdateTime),
		})
	}
	return bus, nil
}

func (ur *udpPacketRepo) GetUDPPackets(ctx context.Context, name string) (*biz.UDPPacket, error) {
	return nil, nil
}

func (ur *udpPacketRepo) CreateUDPPacket(ctx context.Context, udp *biz.UDPPacket) (*biz.UDPPacket, error) {
	_, err := ur.data.db.UDPPacket.Create().Save(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (ur *udpPacketRepo) UpdateUDPPacket(ctx context.Context, udp *biz.UDPPacket) (*biz.UDPPacket, error) {
	return nil, nil
}

func (ur *udpPacketRepo) DeleteUDPPacket(ctx context.Context, name string) error {
	return nil
}
