// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rtc/rtc.proto

package rtc // import "OpenIM/pkg/proto/rtc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sdkws "OpenIM/pkg/proto/sdkws"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignalMessageAssembleReq struct {
	SignalReq            *sdkws.SignalReq `protobuf:"bytes,1,opt,name=signalReq" json:"signalReq,omitempty"`
	OperationID          string           `protobuf:"bytes,2,opt,name=operationID" json:"operationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SignalMessageAssembleReq) Reset()         { *m = SignalMessageAssembleReq{} }
func (m *SignalMessageAssembleReq) String() string { return proto.CompactTextString(m) }
func (*SignalMessageAssembleReq) ProtoMessage()    {}
func (*SignalMessageAssembleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_rtc_62b2557b8d73c35e, []int{0}
}
func (m *SignalMessageAssembleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalMessageAssembleReq.Unmarshal(m, b)
}
func (m *SignalMessageAssembleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalMessageAssembleReq.Marshal(b, m, deterministic)
}
func (dst *SignalMessageAssembleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalMessageAssembleReq.Merge(dst, src)
}
func (m *SignalMessageAssembleReq) XXX_Size() int {
	return xxx_messageInfo_SignalMessageAssembleReq.Size(m)
}
func (m *SignalMessageAssembleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalMessageAssembleReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignalMessageAssembleReq proto.InternalMessageInfo

func (m *SignalMessageAssembleReq) GetSignalReq() *sdkws.SignalReq {
	if m != nil {
		return m.SignalReq
	}
	return nil
}

func (m *SignalMessageAssembleReq) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type SignalMessageAssembleResp struct {
	IsPass               bool              `protobuf:"varint,1,opt,name=isPass" json:"isPass,omitempty"`
	SignalResp           *sdkws.SignalResp `protobuf:"bytes,2,opt,name=signalResp" json:"signalResp,omitempty"`
	MsgData              *sdkws.MsgData    `protobuf:"bytes,3,opt,name=msgData" json:"msgData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SignalMessageAssembleResp) Reset()         { *m = SignalMessageAssembleResp{} }
func (m *SignalMessageAssembleResp) String() string { return proto.CompactTextString(m) }
func (*SignalMessageAssembleResp) ProtoMessage()    {}
func (*SignalMessageAssembleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_rtc_62b2557b8d73c35e, []int{1}
}
func (m *SignalMessageAssembleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalMessageAssembleResp.Unmarshal(m, b)
}
func (m *SignalMessageAssembleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalMessageAssembleResp.Marshal(b, m, deterministic)
}
func (dst *SignalMessageAssembleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalMessageAssembleResp.Merge(dst, src)
}
func (m *SignalMessageAssembleResp) XXX_Size() int {
	return xxx_messageInfo_SignalMessageAssembleResp.Size(m)
}
func (m *SignalMessageAssembleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalMessageAssembleResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignalMessageAssembleResp proto.InternalMessageInfo

func (m *SignalMessageAssembleResp) GetIsPass() bool {
	if m != nil {
		return m.IsPass
	}
	return false
}

func (m *SignalMessageAssembleResp) GetSignalResp() *sdkws.SignalResp {
	if m != nil {
		return m.SignalResp
	}
	return nil
}

func (m *SignalMessageAssembleResp) GetMsgData() *sdkws.MsgData {
	if m != nil {
		return m.MsgData
	}
	return nil
}

type SignalGetRoomsReq struct {
	RoomID               string   `protobuf:"bytes,1,opt,name=roomID" json:"roomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalGetRoomsReq) Reset()         { *m = SignalGetRoomsReq{} }
func (m *SignalGetRoomsReq) String() string { return proto.CompactTextString(m) }
func (*SignalGetRoomsReq) ProtoMessage()    {}
func (*SignalGetRoomsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_rtc_62b2557b8d73c35e, []int{2}
}
func (m *SignalGetRoomsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalGetRoomsReq.Unmarshal(m, b)
}
func (m *SignalGetRoomsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalGetRoomsReq.Marshal(b, m, deterministic)
}
func (dst *SignalGetRoomsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalGetRoomsReq.Merge(dst, src)
}
func (m *SignalGetRoomsReq) XXX_Size() int {
	return xxx_messageInfo_SignalGetRoomsReq.Size(m)
}
func (m *SignalGetRoomsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalGetRoomsReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignalGetRoomsReq proto.InternalMessageInfo

func (m *SignalGetRoomsReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

type SignalGetRoomsResp struct {
	Rooms                []*sdkws.SignalGetRoomByGroupIDReply `protobuf:"bytes,1,rep,name=rooms" json:"rooms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *SignalGetRoomsResp) Reset()         { *m = SignalGetRoomsResp{} }
func (m *SignalGetRoomsResp) String() string { return proto.CompactTextString(m) }
func (*SignalGetRoomsResp) ProtoMessage()    {}
func (*SignalGetRoomsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_rtc_62b2557b8d73c35e, []int{3}
}
func (m *SignalGetRoomsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalGetRoomsResp.Unmarshal(m, b)
}
func (m *SignalGetRoomsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalGetRoomsResp.Marshal(b, m, deterministic)
}
func (dst *SignalGetRoomsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalGetRoomsResp.Merge(dst, src)
}
func (m *SignalGetRoomsResp) XXX_Size() int {
	return xxx_messageInfo_SignalGetRoomsResp.Size(m)
}
func (m *SignalGetRoomsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalGetRoomsResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignalGetRoomsResp proto.InternalMessageInfo

func (m *SignalGetRoomsResp) GetRooms() []*sdkws.SignalGetRoomByGroupIDReply {
	if m != nil {
		return m.Rooms
	}
	return nil
}

func init() {
	proto.RegisterType((*SignalMessageAssembleReq)(nil), "rtc.SignalMessageAssembleReq")
	proto.RegisterType((*SignalMessageAssembleResp)(nil), "rtc.SignalMessageAssembleResp")
	proto.RegisterType((*SignalGetRoomsReq)(nil), "rtc.SignalGetRoomsReq")
	proto.RegisterType((*SignalGetRoomsResp)(nil), "rtc.SignalGetRoomsResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RtcService service

type RtcServiceClient interface {
	SignalMessageAssemble(ctx context.Context, in *SignalMessageAssembleReq, opts ...grpc.CallOption) (*SignalMessageAssembleResp, error)
	SignalGetRooms(ctx context.Context, in *SignalGetRoomsReq, opts ...grpc.CallOption) (*SignalGetRoomsResp, error)
}

type rtcServiceClient struct {
	cc *grpc.ClientConn
}

func NewRtcServiceClient(cc *grpc.ClientConn) RtcServiceClient {
	return &rtcServiceClient{cc}
}

func (c *rtcServiceClient) SignalMessageAssemble(ctx context.Context, in *SignalMessageAssembleReq, opts ...grpc.CallOption) (*SignalMessageAssembleResp, error) {
	out := new(SignalMessageAssembleResp)
	err := grpc.Invoke(ctx, "/rtc.RtcService/SignalMessageAssemble", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) SignalGetRooms(ctx context.Context, in *SignalGetRoomsReq, opts ...grpc.CallOption) (*SignalGetRoomsResp, error) {
	out := new(SignalGetRoomsResp)
	err := grpc.Invoke(ctx, "/rtc.RtcService/SignalGetRooms", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RtcService service

type RtcServiceServer interface {
	SignalMessageAssemble(context.Context, *SignalMessageAssembleReq) (*SignalMessageAssembleResp, error)
	SignalGetRooms(context.Context, *SignalGetRoomsReq) (*SignalGetRoomsResp, error)
}

func RegisterRtcServiceServer(s *grpc.Server, srv RtcServiceServer) {
	s.RegisterService(&_RtcService_serviceDesc, srv)
}

func _RtcService_SignalMessageAssemble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalMessageAssembleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalMessageAssemble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rtc.RtcService/SignalMessageAssemble",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalMessageAssemble(ctx, req.(*SignalMessageAssembleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_SignalGetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalGetRoomsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalGetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rtc.RtcService/SignalGetRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalGetRooms(ctx, req.(*SignalGetRoomsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RtcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rtc.RtcService",
	HandlerType: (*RtcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignalMessageAssemble",
			Handler:    _RtcService_SignalMessageAssemble_Handler,
		},
		{
			MethodName: "SignalGetRooms",
			Handler:    _RtcService_SignalGetRooms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rtc/rtc.proto",
}

func init() { proto.RegisterFile("rtc/rtc.proto", fileDescriptor_rtc_62b2557b8d73c35e) }

var fileDescriptor_rtc_62b2557b8d73c35e = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0x86, 0xf1, 0x86, 0xcd, 0x6e, 0xc6, 0x6c, 0xd8, 0x08, 0x36, 0xeb, 0x1a, 0x5a, 0x82, 0x2f,
	0x35, 0x94, 0xd8, 0xd4, 0xbd, 0x14, 0x7a, 0x4a, 0x30, 0x04, 0x1f, 0xdc, 0x16, 0x05, 0x7a, 0xe8,
	0xcd, 0x71, 0x85, 0x31, 0xb1, 0x23, 0x45, 0xa3, 0x36, 0xe4, 0x2d, 0xfa, 0x16, 0x7d, 0xcd, 0x62,
	0xd9, 0xa1, 0x4e, 0x48, 0x7a, 0x9c, 0xff, 0xff, 0x46, 0xf3, 0x4b, 0x23, 0xf8, 0x23, 0x55, 0xea,
	0x4b, 0x95, 0x7a, 0x42, 0x72, 0xc5, 0x49, 0x47, 0xaa, 0xd4, 0xbe, 0x7c, 0x10, 0x6c, 0x35, 0x8e,
	0xe2, 0xf1, 0x9c, 0xc9, 0x37, 0x26, 0x7d, 0xb1, 0xcc, 0x7c, 0x6d, 0xfb, 0xf8, 0xb2, 0xdc, 0xa0,
	0xbf, 0xc1, 0x9a, 0x76, 0x0a, 0xb0, 0xe6, 0x79, 0xb6, 0x4a, 0x8a, 0x98, 0x21, 0x26, 0x19, 0x9b,
	0x20, 0xb2, 0x72, 0x51, 0x30, 0xca, 0xd6, 0xc4, 0x83, 0x1e, 0x6a, 0x8f, 0xb2, 0xb5, 0x65, 0x8c,
	0x0c, 0xd7, 0x0c, 0xfe, 0x7a, 0xba, 0xdf, 0x9b, 0xef, 0x74, 0xfa, 0x85, 0x90, 0x11, 0x98, 0x5c,
	0x30, 0x99, 0xa8, 0x9c, 0xaf, 0xa2, 0xd0, 0xfa, 0x31, 0x32, 0xdc, 0x1e, 0x6d, 0x4b, 0xce, 0xbb,
	0x01, 0x67, 0x27, 0xc6, 0xa1, 0x20, 0x43, 0xe8, 0xe6, 0xf8, 0x98, 0x20, 0xea, 0x61, 0xbf, 0x69,
	0x53, 0x91, 0x6b, 0x80, 0xdd, 0x10, 0x14, 0xfa, 0x58, 0x33, 0x18, 0x1c, 0x04, 0x41, 0x41, 0x5b,
	0x10, 0x71, 0xe1, 0x57, 0x89, 0x59, 0x98, 0xa8, 0xc4, 0xea, 0x68, 0xbe, 0xdf, 0xf0, 0x71, 0xad,
	0xd2, 0x9d, 0xed, 0x5c, 0xc1, 0xa0, 0x3e, 0x63, 0xc6, 0x14, 0xe5, 0xbc, 0xc4, 0xea, 0x26, 0x43,
	0xe8, 0x4a, 0xce, 0xcb, 0x28, 0xd4, 0x49, 0x7a, 0xb4, 0xa9, 0x9c, 0x7b, 0x20, 0x87, 0x30, 0x0a,
	0x72, 0x0b, 0x3f, 0x2b, 0xbf, 0x8a, 0xdd, 0x71, 0xcd, 0xc0, 0xd9, 0x8b, 0xd6, 0x90, 0xd3, 0xed,
	0x4c, 0xf2, 0x57, 0x11, 0x85, 0x94, 0x89, 0x62, 0x4b, 0xeb, 0x86, 0xe0, 0xc3, 0x00, 0xa0, 0x2a,
	0xad, 0x96, 0x94, 0xa7, 0x8c, 0x3c, 0xc1, 0xbf, 0xa3, 0xaf, 0x43, 0xce, 0xbd, 0x6a, 0xbf, 0xa7,
	0x16, 0x65, 0x5f, 0x7c, 0x67, 0xa3, 0x20, 0x13, 0xe8, 0xef, 0xc7, 0x26, 0xc3, 0x56, 0x47, 0xeb,
	0xe2, 0xf6, 0xff, 0xa3, 0x3a, 0x8a, 0xa9, 0xfd, 0x6c, 0x55, 0x5f, 0x2a, 0x8a, 0x5b, 0x5f, 0x49,
	0xaa, 0xf4, 0x4e, 0xaa, 0x74, 0xd1, 0xd5, 0xe5, 0xcd, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49,
	0x81, 0xcb, 0x8d, 0x89, 0x02, 0x00, 0x00,
}
