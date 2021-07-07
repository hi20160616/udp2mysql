// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/udp2mysql/v1/udp2mysql.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UDPPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the UDPPacket. It must have the format of "shelves/*/books/*".
	// For example: "shelves/shelf1/books/book2".
	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id         string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Title      string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content    string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *UDPPacket) Reset() {
	*x = UDPPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UDPPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UDPPacket) ProtoMessage() {}

func (x *UDPPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UDPPacket.ProtoReflect.Descriptor instead.
func (*UDPPacket) Descriptor() ([]byte, []int) {
	return file_api_udp2mysql_v1_udp2mysql_proto_rawDescGZIP(), []int{0}
}

func (x *UDPPacket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UDPPacket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UDPPacket) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UDPPacket) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UDPPacket) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type ListUDPPacketsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the parent resource where to create the udp_packet.
	// For example: "shelves/shelf1".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListUDPPacketsRequest) Reset() {
	*x = ListUDPPacketsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUDPPacketsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUDPPacketsRequest) ProtoMessage() {}

func (x *ListUDPPacketsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUDPPacketsRequest.ProtoReflect.Descriptor instead.
func (*ListUDPPacketsRequest) Descriptor() ([]byte, []int) {
	return file_api_udp2mysql_v1_udp2mysql_proto_rawDescGZIP(), []int{1}
}

func (x *ListUDPPacketsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListUDPPacketsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListUDPPacketsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListUDPPacketsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The field name should match the noun "udp_packet" in the method name.  There
	// will be a maximum number of items returned based on the page_size field
	// in the request.
	UdpPackets []*UDPPacket `protobuf:"bytes,1,rep,name=udp_packets,json=udpPackets,proto3" json:"udp_packets,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListUDPPacketsResponse) Reset() {
	*x = ListUDPPacketsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUDPPacketsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUDPPacketsResponse) ProtoMessage() {}

func (x *ListUDPPacketsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUDPPacketsResponse.ProtoReflect.Descriptor instead.
func (*ListUDPPacketsResponse) Descriptor() ([]byte, []int) {
	return file_api_udp2mysql_v1_udp2mysql_proto_rawDescGZIP(), []int{2}
}

func (x *ListUDPPacketsResponse) GetUdpPackets() []*UDPPacket {
	if x != nil {
		return x.UdpPackets
	}
	return nil
}

func (x *ListUDPPacketsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetUDPPacketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of a UDPPacket. For example: "shelves/shelf1/books/book2".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetUDPPacketRequest) Reset() {
	*x = GetUDPPacketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUDPPacketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUDPPacketRequest) ProtoMessage() {}

func (x *GetUDPPacketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUDPPacketRequest.ProtoReflect.Descriptor instead.
func (*GetUDPPacketRequest) Descriptor() ([]byte, []int) {
	return file_api_udp2mysql_v1_udp2mysql_proto_rawDescGZIP(), []int{3}
}

func (x *GetUDPPacketRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateUDPPacketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the parent resource where to create the udp_packet.
	// For example: "shelves/shelf1".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The udp_packet id to use for this udp_packet
	UdpPacketId string `protobuf:"bytes,3,opt,name=udp_packet_id,json=udpPacketId,proto3" json:"udp_packet_id,omitempty"`
	// The udp_packet resource to create.
	// The field name should match the Noun in the method name.
	UdpPacket *UDPPacket `protobuf:"bytes,2,opt,name=udp_packet,json=udpPacket,proto3" json:"udp_packet,omitempty"`
}

func (x *CreateUDPPacketRequest) Reset() {
	*x = CreateUDPPacketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUDPPacketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUDPPacketRequest) ProtoMessage() {}

func (x *CreateUDPPacketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUDPPacketRequest.ProtoReflect.Descriptor instead.
func (*CreateUDPPacketRequest) Descriptor() ([]byte, []int) {
	return file_api_udp2mysql_v1_udp2mysql_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUDPPacketRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateUDPPacketRequest) GetUdpPacketId() string {
	if x != nil {
		return x.UdpPacketId
	}
	return ""
}

func (x *CreateUDPPacketRequest) GetUdpPacket() *UDPPacket {
	if x != nil {
		return x.UdpPacket
	}
	return nil
}

type UpdateUDPPacketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the parent resource where to create the udp_packet.
	// For example: "shelves/shelf1".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The udp_packet resource which replaces the resource on the server.
	UdpPacket *UDPPacket `protobuf:"bytes,2,opt,name=udp_packet,json=udpPacket,proto3" json:"udp_packet,omitempty"`
	// The update mask applies to the resource. For the `FieldMask` definition,
	// see https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateUDPPacketRequest) Reset() {
	*x = UpdateUDPPacketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUDPPacketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUDPPacketRequest) ProtoMessage() {}

func (x *UpdateUDPPacketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUDPPacketRequest.ProtoReflect.Descriptor instead.
func (*UpdateUDPPacketRequest) Descriptor() ([]byte, []int) {
	return file_api_udp2mysql_v1_udp2mysql_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUDPPacketRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *UpdateUDPPacketRequest) GetUdpPacket() *UDPPacket {
	if x != nil {
		return x.UdpPacket
	}
	return nil
}

func (x *UpdateUDPPacketRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteUDPPacketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the parent resource where to create the udp_packet.
	// For example: "shelves/shelf1".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteUDPPacketRequest) Reset() {
	*x = DeleteUDPPacketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUDPPacketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUDPPacketRequest) ProtoMessage() {}

func (x *DeleteUDPPacketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUDPPacketRequest.ProtoReflect.Descriptor instead.
func (*DeleteUDPPacketRequest) Descriptor() ([]byte, []int) {
	return file_api_udp2mysql_v1_udp2mysql_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteUDPPacketRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *DeleteUDPPacketRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_udp2mysql_v1_udp2mysql_proto protoreflect.FileDescriptor

var file_api_udp2mysql_v1_udp2mysql_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9c, 0x01, 0x0a, 0x09, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x6b, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7a, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x64,
	0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x44, 0x50, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x0a, 0x75, 0x64, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55,
	0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x44,
	0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x64, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0a, 0x75, 0x64,
	0x70, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x44,
	0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x09, 0x75, 0x64, 0x70, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x44, 0x50,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x64, 0x70, 0x32,
	0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x09, 0x75, 0x64, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x44, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0xb6, 0x03, 0x0a, 0x0c, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x70,
	0x69, 0x12, 0x5d, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x64, 0x70, 0x32, 0x6d,
	0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x44, 0x50, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x21, 0x2e, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x24, 0x2e, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79,
	0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x22, 0x00, 0x12, 0x52, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x44, 0x50, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x44, 0x50, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x64,
	0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x44, 0x50, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x75, 0x64, 0x70, 0x32,
	0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x44, 0x50, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x32, 0x30, 0x31, 0x36, 0x30, 0x36,
	0x31, 0x36, 0x2f, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x75, 0x64, 0x70, 0x32, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_udp2mysql_v1_udp2mysql_proto_rawDescOnce sync.Once
	file_api_udp2mysql_v1_udp2mysql_proto_rawDescData = file_api_udp2mysql_v1_udp2mysql_proto_rawDesc
)

func file_api_udp2mysql_v1_udp2mysql_proto_rawDescGZIP() []byte {
	file_api_udp2mysql_v1_udp2mysql_proto_rawDescOnce.Do(func() {
		file_api_udp2mysql_v1_udp2mysql_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_udp2mysql_v1_udp2mysql_proto_rawDescData)
	})
	return file_api_udp2mysql_v1_udp2mysql_proto_rawDescData
}

var file_api_udp2mysql_v1_udp2mysql_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_udp2mysql_v1_udp2mysql_proto_goTypes = []interface{}{
	(*UDPPacket)(nil),              // 0: udp2mysql.v1.UDPPacket
	(*ListUDPPacketsRequest)(nil),  // 1: udp2mysql.v1.ListUDPPacketsRequest
	(*ListUDPPacketsResponse)(nil), // 2: udp2mysql.v1.ListUDPPacketsResponse
	(*GetUDPPacketRequest)(nil),    // 3: udp2mysql.v1.GetUDPPacketRequest
	(*CreateUDPPacketRequest)(nil), // 4: udp2mysql.v1.CreateUDPPacketRequest
	(*UpdateUDPPacketRequest)(nil), // 5: udp2mysql.v1.UpdateUDPPacketRequest
	(*DeleteUDPPacketRequest)(nil), // 6: udp2mysql.v1.DeleteUDPPacketRequest
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil),  // 8: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_api_udp2mysql_v1_udp2mysql_proto_depIdxs = []int32{
	7,  // 0: udp2mysql.v1.UDPPacket.update_time:type_name -> google.protobuf.Timestamp
	0,  // 1: udp2mysql.v1.ListUDPPacketsResponse.udp_packets:type_name -> udp2mysql.v1.UDPPacket
	0,  // 2: udp2mysql.v1.CreateUDPPacketRequest.udp_packet:type_name -> udp2mysql.v1.UDPPacket
	0,  // 3: udp2mysql.v1.UpdateUDPPacketRequest.udp_packet:type_name -> udp2mysql.v1.UDPPacket
	8,  // 4: udp2mysql.v1.UpdateUDPPacketRequest.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 5: udp2mysql.v1.UDPPacketApi.ListUDPPackets:input_type -> udp2mysql.v1.ListUDPPacketsRequest
	3,  // 6: udp2mysql.v1.UDPPacketApi.GetUDPPacket:input_type -> udp2mysql.v1.GetUDPPacketRequest
	4,  // 7: udp2mysql.v1.UDPPacketApi.CreateUDPPacket:input_type -> udp2mysql.v1.CreateUDPPacketRequest
	5,  // 8: udp2mysql.v1.UDPPacketApi.UpdateUDPPacket:input_type -> udp2mysql.v1.UpdateUDPPacketRequest
	6,  // 9: udp2mysql.v1.UDPPacketApi.DeleteUDPPacket:input_type -> udp2mysql.v1.DeleteUDPPacketRequest
	2,  // 10: udp2mysql.v1.UDPPacketApi.ListUDPPackets:output_type -> udp2mysql.v1.ListUDPPacketsResponse
	0,  // 11: udp2mysql.v1.UDPPacketApi.GetUDPPacket:output_type -> udp2mysql.v1.UDPPacket
	0,  // 12: udp2mysql.v1.UDPPacketApi.CreateUDPPacket:output_type -> udp2mysql.v1.UDPPacket
	0,  // 13: udp2mysql.v1.UDPPacketApi.UpdateUDPPacket:output_type -> udp2mysql.v1.UDPPacket
	9,  // 14: udp2mysql.v1.UDPPacketApi.DeleteUDPPacket:output_type -> google.protobuf.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_udp2mysql_v1_udp2mysql_proto_init() }
func file_api_udp2mysql_v1_udp2mysql_proto_init() {
	if File_api_udp2mysql_v1_udp2mysql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UDPPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUDPPacketsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUDPPacketsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUDPPacketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUDPPacketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUDPPacketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_udp2mysql_v1_udp2mysql_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUDPPacketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_udp2mysql_v1_udp2mysql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_udp2mysql_v1_udp2mysql_proto_goTypes,
		DependencyIndexes: file_api_udp2mysql_v1_udp2mysql_proto_depIdxs,
		MessageInfos:      file_api_udp2mysql_v1_udp2mysql_proto_msgTypes,
	}.Build()
	File_api_udp2mysql_v1_udp2mysql_proto = out.File
	file_api_udp2mysql_v1_udp2mysql_proto_rawDesc = nil
	file_api_udp2mysql_v1_udp2mysql_proto_goTypes = nil
	file_api_udp2mysql_v1_udp2mysql_proto_depIdxs = nil
}
