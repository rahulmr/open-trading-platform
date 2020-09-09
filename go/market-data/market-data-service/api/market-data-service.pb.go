// Code generated by protoc-gen-go. DO NOT EDIT.
// source: market-data-service.proto

package api

import (
	"github.com/ettec/otp-common/model"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MdsConnectRequest struct {
	SubscriberId         string   `protobuf:"bytes,1,opt,name=subscriberId,proto3" json:"subscriberId,omitempty"`
	MaxQuotePerSecond    int32    `protobuf:"varint,2,opt,name=maxQuotePerSecond,proto3" json:"maxQuotePerSecond,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MdsConnectRequest) Reset()         { *m = MdsConnectRequest{} }
func (m *MdsConnectRequest) String() string { return proto.CompactTextString(m) }
func (*MdsConnectRequest) ProtoMessage()    {}
func (*MdsConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d3bea125c825790, []int{0}
}

func (m *MdsConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdsConnectRequest.Unmarshal(m, b)
}
func (m *MdsConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdsConnectRequest.Marshal(b, m, deterministic)
}
func (m *MdsConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdsConnectRequest.Merge(m, src)
}
func (m *MdsConnectRequest) XXX_Size() int {
	return xxx_messageInfo_MdsConnectRequest.Size(m)
}
func (m *MdsConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MdsConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MdsConnectRequest proto.InternalMessageInfo

func (m *MdsConnectRequest) GetSubscriberId() string {
	if m != nil {
		return m.SubscriberId
	}
	return ""
}

func (m *MdsConnectRequest) GetMaxQuotePerSecond() int32 {
	if m != nil {
		return m.MaxQuotePerSecond
	}
	return 0
}

type MdsSubscribeRequest struct {
	SubscriberId         string   `protobuf:"bytes,1,opt,name=subscriberId,proto3" json:"subscriberId,omitempty"`
	ListingId            int32    `protobuf:"varint,2,opt,name=listingId,proto3" json:"listingId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MdsSubscribeRequest) Reset()         { *m = MdsSubscribeRequest{} }
func (m *MdsSubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*MdsSubscribeRequest) ProtoMessage()    {}
func (*MdsSubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d3bea125c825790, []int{1}
}

func (m *MdsSubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MdsSubscribeRequest.Unmarshal(m, b)
}
func (m *MdsSubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MdsSubscribeRequest.Marshal(b, m, deterministic)
}
func (m *MdsSubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MdsSubscribeRequest.Merge(m, src)
}
func (m *MdsSubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_MdsSubscribeRequest.Size(m)
}
func (m *MdsSubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MdsSubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MdsSubscribeRequest proto.InternalMessageInfo

func (m *MdsSubscribeRequest) GetSubscriberId() string {
	if m != nil {
		return m.SubscriberId
	}
	return ""
}

func (m *MdsSubscribeRequest) GetListingId() int32 {
	if m != nil {
		return m.ListingId
	}
	return 0
}

func init() {
	proto.RegisterType((*MdsConnectRequest)(nil), "marketdataservice.MdsConnectRequest")
	proto.RegisterType((*MdsSubscribeRequest)(nil), "marketdataservice.MdsSubscribeRequest")
}

func init() { proto.RegisterFile("market-data-service.proto", fileDescriptor_2d3bea125c825790) }

var fileDescriptor_2d3bea125c825790 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x9b, 0xef, 0xa3, 0x4a, 0x86, 0x82, 0x66, 0xdc, 0xd4, 0xe0, 0xa2, 0x04, 0x91, 0x2e,
	0xec, 0x20, 0xfa, 0x06, 0xc6, 0x2e, 0xba, 0x08, 0x68, 0xb2, 0x10, 0xdc, 0xcd, 0x9f, 0x8b, 0x0c,
	0x66, 0xe6, 0xa6, 0x33, 0x13, 0xd1, 0xa7, 0xf1, 0x55, 0xa5, 0x49, 0xaa, 0x94, 0x76, 0xe3, 0xf6,
	0x9c, 0x99, 0x73, 0xee, 0xf9, 0x91, 0x73, 0xc3, 0xdd, 0x1b, 0x84, 0x85, 0xe2, 0x81, 0x2f, 0x3c,
	0xb8, 0x77, 0x2d, 0x81, 0x35, 0x0e, 0x03, 0xd2, 0xa4, 0xb7, 0x36, 0xce, 0x60, 0xa4, 0x89, 0x41,
	0x05, 0xb5, 0x44, 0x63, 0xd0, 0xf6, 0xaf, 0xd2, 0x13, 0x59, 0xa3, 0x58, 0xb7, 0x18, 0x86, 0x6f,
	0x19, 0x90, 0xa4, 0x50, 0x3e, 0x47, 0x6b, 0x41, 0x86, 0x12, 0xd6, 0x2d, 0xf8, 0x40, 0x33, 0x32,
	0xf1, 0xad, 0xf0, 0xd2, 0x69, 0x01, 0x6e, 0xa5, 0xa6, 0xd1, 0x2c, 0x9a, 0xc7, 0xe5, 0x8e, 0x46,
	0xaf, 0x49, 0x62, 0xf8, 0xc7, 0xd3, 0x26, 0xea, 0x11, 0x5c, 0x05, 0x12, 0xad, 0x9a, 0xfe, 0x9b,
	0x45, 0xf3, 0x71, 0xb9, 0x6f, 0x64, 0xcf, 0xe4, 0xac, 0x50, 0xbe, 0xda, 0x06, 0xfc, 0xa5, 0xe8,
	0x82, 0xc4, 0xb5, 0xf6, 0x41, 0xdb, 0xd7, 0xd5, 0xb6, 0xe0, 0x57, 0xb8, 0xfd, 0x8a, 0x48, 0x52,
	0x74, 0xcb, 0x1f, 0x78, 0xe0, 0x55, 0xbf, 0x9c, 0xe6, 0x24, 0xfe, 0xe9, 0xa2, 0x57, 0x6c, 0x0f,
	0x0d, 0x3b, 0x70, 0x4c, 0x3a, 0x61, 0x1d, 0x2f, 0xb6, 0x34, 0x4d, 0xf8, 0xcc, 0x46, 0x74, 0x49,
	0x8e, 0x07, 0x2e, 0xf4, 0xf2, 0x70, 0xc4, 0x2e, 0xb6, 0xf4, 0x74, 0x08, 0xc8, 0x6b, 0x14, 0xdd,
	0xfe, 0x6c, 0x74, 0x13, 0xdd, 0x8f, 0x5f, 0xfe, 0xf3, 0x46, 0x8b, 0xa3, 0x8e, 0xf7, 0xdd, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x42, 0xe4, 0x86, 0xc3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MarketDataServiceClient is the client API for MarketDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MarketDataServiceClient interface {
	Subscribe(ctx context.Context, in *MdsSubscribeRequest, opts ...grpc.CallOption) (*model.Empty, error)
	Connect(ctx context.Context, in *MdsConnectRequest, opts ...grpc.CallOption) (MarketDataService_ConnectClient, error)
}

type marketDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewMarketDataServiceClient(cc *grpc.ClientConn) MarketDataServiceClient {
	return &marketDataServiceClient{cc}
}

func (c *marketDataServiceClient) Subscribe(ctx context.Context, in *MdsSubscribeRequest, opts ...grpc.CallOption) (*model.Empty, error) {
	out := new(model.Empty)
	err := c.cc.Invoke(ctx, "/marketdataservice.MarketDataService/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) Connect(ctx context.Context, in *MdsConnectRequest, opts ...grpc.CallOption) (MarketDataService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MarketDataService_serviceDesc.Streams[0], "/marketdataservice.MarketDataService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketDataServiceConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketDataService_ConnectClient interface {
	Recv() (*model.ClobQuote, error)
	grpc.ClientStream
}

type marketDataServiceConnectClient struct {
	grpc.ClientStream
}

func (x *marketDataServiceConnectClient) Recv() (*model.ClobQuote, error) {
	m := new(model.ClobQuote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MarketDataServiceServer is the server API for MarketDataService service.
type MarketDataServiceServer interface {
	Subscribe(context.Context, *MdsSubscribeRequest) (*model.Empty, error)
	Connect(*MdsConnectRequest, MarketDataService_ConnectServer) error
}

// UnimplementedMarketDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMarketDataServiceServer struct {
}

func (*UnimplementedMarketDataServiceServer) Subscribe(ctx context.Context, req *MdsSubscribeRequest) (*model.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (*UnimplementedMarketDataServiceServer) Connect(req *MdsConnectRequest, srv MarketDataService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterMarketDataServiceServer(s *grpc.Server, srv MarketDataServiceServer) {
	s.RegisterService(&_MarketDataService_serviceDesc, srv)
}

func _MarketDataService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MdsSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketdataservice.MarketDataService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).Subscribe(ctx, req.(*MdsSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MdsConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketDataServiceServer).Connect(m, &marketDataServiceConnectServer{stream})
}

type MarketDataService_ConnectServer interface {
	Send(*model.ClobQuote) error
	grpc.ServerStream
}

type marketDataServiceConnectServer struct {
	grpc.ServerStream
}

func (x *marketDataServiceConnectServer) Send(m *model.ClobQuote) error {
	return x.ServerStream.SendMsg(m)
}

var _MarketDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "marketdataservice.MarketDataService",
	HandlerType: (*MarketDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _MarketDataService_Subscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _MarketDataService_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "market-data-service.proto",
}
