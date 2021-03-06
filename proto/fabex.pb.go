// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fabex.proto

package proto

import (
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

type RequestRange struct {
	Startblock           int64    `protobuf:"varint,1,opt,name=startblock,proto3" json:"startblock,omitempty"`
	Endblock             int64    `protobuf:"varint,2,opt,name=endblock,proto3" json:"endblock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestRange) Reset()         { *m = RequestRange{} }
func (m *RequestRange) String() string { return proto.CompactTextString(m) }
func (*RequestRange) ProtoMessage()    {}
func (*RequestRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d7206373264ff4, []int{0}
}

func (m *RequestRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestRange.Unmarshal(m, b)
}
func (m *RequestRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestRange.Marshal(b, m, deterministic)
}
func (m *RequestRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestRange.Merge(m, src)
}
func (m *RequestRange) XXX_Size() int {
	return xxx_messageInfo_RequestRange.Size(m)
}
func (m *RequestRange) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestRange.DiscardUnknown(m)
}

var xxx_messageInfo_RequestRange proto.InternalMessageInfo

func (m *RequestRange) GetStartblock() int64 {
	if m != nil {
		return m.Startblock
	}
	return 0
}

func (m *RequestRange) GetEndblock() int64 {
	if m != nil {
		return m.Endblock
	}
	return 0
}

type Entry struct {
	Channelid            string   `protobuf:"bytes,1,opt,name=channelid,proto3" json:"channelid,omitempty"`
	Txid                 string   `protobuf:"bytes,2,opt,name=txid,proto3" json:"txid,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Previoushash         string   `protobuf:"bytes,4,opt,name=previoushash,proto3" json:"previoushash,omitempty"`
	Blocknum             uint64   `protobuf:"varint,5,opt,name=blocknum,proto3" json:"blocknum,omitempty"`
	Payload              string   `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	Time                 int64    `protobuf:"varint,7,opt,name=time,proto3" json:"time,omitempty"`
	Validationcode       int32    `protobuf:"varint,8,opt,name=validationcode,proto3" json:"validationcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d7206373264ff4, []int{1}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetChannelid() string {
	if m != nil {
		return m.Channelid
	}
	return ""
}

func (m *Entry) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *Entry) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Entry) GetPrevioushash() string {
	if m != nil {
		return m.Previoushash
	}
	return ""
}

func (m *Entry) GetBlocknum() uint64 {
	if m != nil {
		return m.Blocknum
	}
	return 0
}

func (m *Entry) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *Entry) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Entry) GetValidationcode() int32 {
	if m != nil {
		return m.Validationcode
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestRange)(nil), "fabex.RequestRange")
	proto.RegisterType((*Entry)(nil), "fabex.Entry")
}

func init() { proto.RegisterFile("fabex.proto", fileDescriptor_d7d7206373264ff4) }

var fileDescriptor_d7d7206373264ff4 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x4e, 0xeb, 0x30,
	0x10, 0x85, 0x95, 0xdb, 0xa6, 0x3f, 0x73, 0x2b, 0x16, 0xc3, 0xc6, 0xaa, 0x10, 0xaa, 0x8a, 0x04,
	0x65, 0x41, 0x83, 0xe0, 0x0d, 0x90, 0xa0, 0x12, 0x4b, 0x2f, 0x91, 0x58, 0x38, 0xf1, 0xd0, 0x58,
	0xa4, 0x76, 0x70, 0x9c, 0xaa, 0x7d, 0x5e, 0x5e, 0x04, 0x65, 0x52, 0x4a, 0xdb, 0x55, 0xe6, 0x7c,
	0x27, 0x63, 0xcf, 0x1c, 0xc3, 0xff, 0x0f, 0x95, 0xd2, 0x66, 0x5e, 0x7a, 0x17, 0x1c, 0xc6, 0x2c,
	0xa6, 0xaf, 0x30, 0x92, 0xf4, 0x55, 0x53, 0x15, 0xa4, 0xb2, 0x4b, 0xc2, 0x4b, 0x80, 0x2a, 0x28,
	0x1f, 0xd2, 0xc2, 0x65, 0x9f, 0x22, 0x9a, 0x44, 0xb3, 0x8e, 0x3c, 0x20, 0x38, 0x86, 0x01, 0x59,
	0xdd, 0xba, 0xff, 0xd8, 0xdd, 0xeb, 0xe9, 0x77, 0x04, 0xf1, 0xb3, 0x0d, 0x7e, 0x8b, 0x17, 0x30,
	0xcc, 0x72, 0x65, 0x2d, 0x15, 0x46, 0xf3, 0x21, 0x43, 0xf9, 0x07, 0x10, 0xa1, 0x1b, 0x36, 0x46,
	0x73, 0xff, 0x50, 0x72, 0xdd, 0xb0, 0x5c, 0x55, 0xb9, 0xe8, 0xb4, 0xac, 0xa9, 0x71, 0x0a, 0xa3,
	0xd2, 0xd3, 0xda, 0xb8, 0xba, 0x62, 0xaf, 0xcb, 0xde, 0x11, 0x6b, 0xe6, 0xe1, 0xcb, 0x6d, 0xbd,
	0x12, 0xf1, 0x24, 0x9a, 0x75, 0xe5, 0x5e, 0xa3, 0x80, 0x7e, 0xa9, 0xb6, 0x85, 0x53, 0x5a, 0xf4,
	0xb8, 0xf5, 0x57, 0xf2, 0x04, 0x66, 0x45, 0xa2, 0xcf, 0x1b, 0x70, 0x8d, 0xd7, 0x70, 0xb6, 0x56,
	0x85, 0xd1, 0x2a, 0x18, 0x67, 0x33, 0xa7, 0x49, 0x0c, 0x26, 0xd1, 0x2c, 0x96, 0x27, 0xf4, 0xe1,
	0x1d, 0xe2, 0x97, 0x26, 0x3a, 0xbc, 0x82, 0xce, 0x82, 0x02, 0x8e, 0xe6, 0x6d, 0xac, 0xbc, 0xf9,
	0xf8, 0x48, 0xdd, 0x47, 0x98, 0xc0, 0x60, 0x41, 0xbb, 0x6c, 0xcf, 0x77, 0xde, 0x61, 0xe0, 0xa7,
	0x0d, 0x4f, 0xb7, 0x6f, 0x37, 0x4b, 0x13, 0xf2, 0x3a, 0x9d, 0x67, 0x6e, 0x95, 0xe4, 0xdb, 0x92,
	0x7c, 0x41, 0x7a, 0x49, 0xfe, 0xae, 0x50, 0x69, 0x95, 0xf0, 0xcf, 0x09, 0x3f, 0x61, 0xda, 0xe3,
	0xcf, 0xe3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0xc0, 0x76, 0x0a, 0xd8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FabexClient is the client API for Fabex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FabexClient interface {
	Get(ctx context.Context, in *Entry, opts ...grpc.CallOption) (Fabex_GetClient, error)
	GetRange(ctx context.Context, in *RequestRange, opts ...grpc.CallOption) (Fabex_GetRangeClient, error)
}

type fabexClient struct {
	cc *grpc.ClientConn
}

func NewFabexClient(cc *grpc.ClientConn) FabexClient {
	return &fabexClient{cc}
}

func (c *fabexClient) Get(ctx context.Context, in *Entry, opts ...grpc.CallOption) (Fabex_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[0], "/fabex.Fabex/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_GetClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type fabexGetClient struct {
	grpc.ClientStream
}

func (x *fabexGetClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fabexClient) GetRange(ctx context.Context, in *RequestRange, opts ...grpc.CallOption) (Fabex_GetRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Fabex_serviceDesc.Streams[1], "/fabex.Fabex/GetRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &fabexGetRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fabex_GetRangeClient interface {
	Recv() (*Entry, error)
	grpc.ClientStream
}

type fabexGetRangeClient struct {
	grpc.ClientStream
}

func (x *fabexGetRangeClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FabexServer is the server API for Fabex service.
type FabexServer interface {
	Get(*Entry, Fabex_GetServer) error
	GetRange(*RequestRange, Fabex_GetRangeServer) error
}

// UnimplementedFabexServer can be embedded to have forward compatible implementations.
type UnimplementedFabexServer struct {
}

func (*UnimplementedFabexServer) Get(req *Entry, srv Fabex_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedFabexServer) GetRange(req *RequestRange, srv Fabex_GetRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRange not implemented")
}

func RegisterFabexServer(s *grpc.Server, srv FabexServer) {
	s.RegisterService(&_Fabex_serviceDesc, srv)
}

func _Fabex_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Entry)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).Get(m, &fabexGetServer{stream})
}

type Fabex_GetServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type fabexGetServer struct {
	grpc.ServerStream
}

func (x *fabexGetServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

func _Fabex_GetRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestRange)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FabexServer).GetRange(m, &fabexGetRangeServer{stream})
}

type Fabex_GetRangeServer interface {
	Send(*Entry) error
	grpc.ServerStream
}

type fabexGetRangeServer struct {
	grpc.ServerStream
}

func (x *fabexGetRangeServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

var _Fabex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fabex.Fabex",
	HandlerType: (*FabexServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _Fabex_Get_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetRange",
			Handler:       _Fabex_GetRange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fabex.proto",
}
