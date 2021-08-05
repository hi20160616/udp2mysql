package data

import (
	"context"

	"github.com/hi20160616/udp2mysql/internal/biz"
	"github.com/hi20160616/udp2mysql/internal/data/db/mariadb"
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

func (ur *udpPacketRepo) ListUDPPackets(ctx context.Context) (*biz.UDPPackets, error) {
	us, err := ur.data.dbClient.UDPPacket.Query().All(ctx)
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
	return &biz.UDPPackets{
		UdpPackets: bus,
	}, nil
}

func (ur *udpPacketRepo) GetUDPPacket(ctx context.Context, name string) (*biz.UDPPacket, error) {
	upkt, err := ur.data.dbClient.UDPPacket.Query().Where([4]string{"name", "=", name}).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.UDPPacket{
		Id:         upkt.ID,
		Name:       upkt.Name,
		Title:      upkt.Title,
		Content:    upkt.Content,
		UpdateTime: timestamppb.New(upkt.UpdateTime),
	}, nil
}

func (ur *udpPacketRepo) CreateUDPPacket(ctx context.Context, udp *biz.UDPPacket) (*biz.UDPPacket, error) {
	upktEnt := &mariadb.UDPPacket{
		ID:         udp.Id,
		Name:       udp.Name,
		Title:      udp.Title,
		Content:    udp.Content,
		UpdateTime: udp.UpdateTime.AsTime(),
	}
	if err := ur.data.dbClient.UDPPacket.Insert(ctx, upktEnt); err != nil {
		return nil, err
	}
	return udp, nil
}

func (ur *udpPacketRepo) UpdateUDPPacket(ctx context.Context, udp *biz.UDPPacket) (*biz.UDPPacket, error) {
	upktEnt := &mariadb.UDPPacket{
		ID:         udp.Id,
		Name:       udp.Name,
		Title:      udp.Title,
		Content:    udp.Content,
		UpdateTime: udp.UpdateTime.AsTime(),
	}
	if err := ur.data.dbClient.UDPPacket.Update(ctx, upktEnt); err != nil {
		return nil, err
	}
	return udp, nil
}

func (ur *udpPacketRepo) DeleteUDPPacket(ctx context.Context, name string) error {
	return ur.data.dbClient.UDPPacket.Delete(ctx, name)
}
