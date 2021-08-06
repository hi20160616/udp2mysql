package service

import (
	"context"

	pb "github.com/hi20160616/udp2mysql/api/udp2mysql/v1"
	"github.com/hi20160616/udp2mysql/internal/biz"
	"github.com/hi20160616/udp2mysql/internal/data"
	"github.com/hi20160616/udp2mysql/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UDPService struct {
	pb.UnimplementedUDPPacketApiServer
	udp *biz.UDPPacketUsecase
}

func InitUDPService() *UDPService {
	dbc := mariadb.NewClient()
	db := &data.Data{DBClient: dbc}
	repo := data.NewUDPPacketRepo(db)
	udpUsecase := biz.NewUDPPacketUsecase(repo)
	return &UDPService{udp: udpUsecase}
}

func (us *UDPService) ListUDPPackets(ctx context.Context, req *pb.ListUDPPacketsRequest) (*pb.ListUDPPacketsResponse, error) {
	udps, err := us.udp.List(ctx)
	if err != nil {
		return nil, err
	}
	rt := &pb.ListUDPPacketsResponse{}
	for _, e := range udps.UdpPackets {
		rt.UdpPackets = append(rt.UdpPackets, &pb.UDPPacket{
			Id:         e.Id,
			Name:       e.Name,
			Title:      e.Title,
			Content:    e.Content,
			UpdateTime: e.UpdateTime,
		})
	}
	return rt, nil
}

func (us *UDPService) GetUDPPacket(ctx context.Context, req *pb.GetUDPPacketRequest) (*pb.UDPPacket, error) {
	udp, err := us.udp.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.UDPPacket{
		Id:         udp.Id,
		Name:       udp.Name,
		Title:      udp.Title,
		Content:    udp.Content,
		UpdateTime: udp.UpdateTime,
	}, nil
}

func (us *UDPService) CreateUDPPacket(ctx context.Context, req *pb.CreateUDPPacketRequest) (*pb.UDPPacket, error) {
	udp, err := us.udp.Create(ctx, &biz.UDPPacket{
		Id:         req.UdpPacket.Id,
		Name:       req.UdpPacket.Name,
		Title:      req.UdpPacket.Title,
		Content:    req.UdpPacket.Content,
		UpdateTime: req.UdpPacket.UpdateTime,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UDPPacket{
		Id:         udp.Id,
		Name:       udp.Name,
		Title:      udp.Title,
		Content:    udp.Content,
		UpdateTime: udp.UpdateTime,
	}, nil
}

func (us *UDPService) UpdateUDPPacket(ctx context.Context, req *pb.UpdateUDPPacketRequest) (*pb.UDPPacket, error) {
	udp, err := us.udp.Update(ctx, &biz.UDPPacket{
		Id:         req.UdpPacket.Id,
		Name:       req.UdpPacket.Name,
		Title:      req.UdpPacket.Title,
		Content:    req.UdpPacket.Content,
		UpdateTime: req.UdpPacket.UpdateTime,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UDPPacket{
		Id:         udp.Id,
		Name:       udp.Name,
		Title:      udp.Title,
		Content:    udp.Content,
		UpdateTime: udp.UpdateTime,
	}, nil
}

func (us *UDPService) DeleteUDPPacket(ctx context.Context, req *pb.DeleteUDPPacketRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, us.udp.Delete(ctx, req.Name)
}
