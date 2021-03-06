// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cns/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgRegisterChainName struct {
	ChainName          string       `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	Owner              string       `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Seed               []string     `protobuf:"bytes,3,rep,name=seed,proto3" json:"seed,omitempty"`
	SourceCodeUrl      string       `protobuf:"bytes,4,opt,name=source_code_url,json=sourceCodeUrl,proto3" json:"source_code_url,omitempty"`
	CanonicalIbcClient string       `protobuf:"bytes,5,opt,name=canonical_ibc_client,json=canonicalIbcClient,proto3" json:"canonical_ibc_client,omitempty"`
	Version            *VersionInfo `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *MsgRegisterChainName) Reset()         { *m = MsgRegisterChainName{} }
func (m *MsgRegisterChainName) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterChainName) ProtoMessage()    {}
func (*MsgRegisterChainName) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{0}
}
func (m *MsgRegisterChainName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterChainName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterChainName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterChainName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterChainName.Merge(m, src)
}
func (m *MsgRegisterChainName) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterChainName) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterChainName.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterChainName proto.InternalMessageInfo

func (m *MsgRegisterChainName) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *MsgRegisterChainName) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgRegisterChainName) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *MsgRegisterChainName) GetSourceCodeUrl() string {
	if m != nil {
		return m.SourceCodeUrl
	}
	return ""
}

func (m *MsgRegisterChainName) GetCanonicalIbcClient() string {
	if m != nil {
		return m.CanonicalIbcClient
	}
	return ""
}

func (m *MsgRegisterChainName) GetVersion() *VersionInfo {
	if m != nil {
		return m.Version
	}
	return nil
}

type MsgRegisterChainNameResponse struct {
}

func (m *MsgRegisterChainNameResponse) Reset()         { *m = MsgRegisterChainNameResponse{} }
func (m *MsgRegisterChainNameResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterChainNameResponse) ProtoMessage()    {}
func (*MsgRegisterChainNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{1}
}
func (m *MsgRegisterChainNameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterChainNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterChainNameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterChainNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterChainNameResponse.Merge(m, src)
}
func (m *MsgRegisterChainNameResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterChainNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterChainNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterChainNameResponse proto.InternalMessageInfo

type MsgUpdateChainInfo struct {
	ChainName          string       `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	Owner              string       `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Seed               []string     `protobuf:"bytes,3,rep,name=seed,proto3" json:"seed,omitempty"`
	SourceCodeUrl      string       `protobuf:"bytes,4,opt,name=source_code_url,json=sourceCodeUrl,proto3" json:"source_code_url,omitempty"`
	CanonicalIbcClient string       `protobuf:"bytes,5,opt,name=canonical_ibc_client,json=canonicalIbcClient,proto3" json:"canonical_ibc_client,omitempty"`
	Version            *VersionInfo `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *MsgUpdateChainInfo) Reset()         { *m = MsgUpdateChainInfo{} }
func (m *MsgUpdateChainInfo) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateChainInfo) ProtoMessage()    {}
func (*MsgUpdateChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{2}
}
func (m *MsgUpdateChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateChainInfo.Merge(m, src)
}
func (m *MsgUpdateChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateChainInfo proto.InternalMessageInfo

func (m *MsgUpdateChainInfo) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *MsgUpdateChainInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgUpdateChainInfo) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *MsgUpdateChainInfo) GetSourceCodeUrl() string {
	if m != nil {
		return m.SourceCodeUrl
	}
	return ""
}

func (m *MsgUpdateChainInfo) GetCanonicalIbcClient() string {
	if m != nil {
		return m.CanonicalIbcClient
	}
	return ""
}

func (m *MsgUpdateChainInfo) GetVersion() *VersionInfo {
	if m != nil {
		return m.Version
	}
	return nil
}

type MsgUpdateChainInfoResponse struct {
}

func (m *MsgUpdateChainInfoResponse) Reset()         { *m = MsgUpdateChainInfoResponse{} }
func (m *MsgUpdateChainInfoResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateChainInfoResponse) ProtoMessage()    {}
func (*MsgUpdateChainInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6219fe2c381986e6, []int{3}
}
func (m *MsgUpdateChainInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateChainInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateChainInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateChainInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateChainInfoResponse.Merge(m, src)
}
func (m *MsgUpdateChainInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateChainInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateChainInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateChainInfoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterChainName)(nil), "tendermint.cns.cns.MsgRegisterChainName")
	proto.RegisterType((*MsgRegisterChainNameResponse)(nil), "tendermint.cns.cns.MsgRegisterChainNameResponse")
	proto.RegisterType((*MsgUpdateChainInfo)(nil), "tendermint.cns.cns.MsgUpdateChainInfo")
	proto.RegisterType((*MsgUpdateChainInfoResponse)(nil), "tendermint.cns.cns.MsgUpdateChainInfoResponse")
}

func init() { proto.RegisterFile("cns/tx.proto", fileDescriptor_6219fe2c381986e6) }

var fileDescriptor_6219fe2c381986e6 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x93, 0xbf, 0xaa, 0xdb, 0x30,
	0x14, 0xc6, 0xa3, 0xe6, 0x4f, 0x89, 0xda, 0x10, 0x2a, 0x3c, 0x18, 0x93, 0xba, 0x21, 0xd0, 0xe0,
	0xc9, 0x0e, 0xe9, 0xd4, 0xa9, 0xd0, 0x4c, 0x19, 0xd2, 0xc1, 0x90, 0x0e, 0x5d, 0x8c, 0x2d, 0x9f,
	0x3a, 0x02, 0x5b, 0x32, 0x92, 0xd2, 0xa6, 0x6f, 0xd1, 0xc7, 0xea, 0x98, 0xb1, 0x63, 0x9b, 0xbc,
	0xc1, 0xe5, 0x3e, 0xc0, 0xc5, 0x32, 0xce, 0x85, 0x24, 0x17, 0x32, 0xdf, 0xc1, 0x46, 0xc7, 0xdf,
	0xef, 0x9c, 0xc3, 0xf9, 0xac, 0x83, 0x5f, 0x53, 0xae, 0x02, 0xbd, 0xf3, 0x4b, 0x29, 0xb4, 0x20,
	0x44, 0x03, 0x4f, 0x41, 0x16, 0x8c, 0x6b, 0x9f, 0x72, 0x55, 0x3d, 0x8e, 0x95, 0x89, 0x4c, 0x18,
	0x39, 0xa8, 0x4e, 0x35, 0xe9, 0x0c, 0xaa, 0xbc, 0x0a, 0x31, 0xe1, 0xe4, 0x1e, 0x61, 0x6b, 0xa5,
	0xb2, 0x10, 0x32, 0xa6, 0x34, 0xc8, 0xc5, 0x26, 0x66, 0xfc, 0x4b, 0x5c, 0x00, 0x79, 0x8b, 0x31,
	0xad, 0x82, 0x88, 0xc7, 0x05, 0xd8, 0x68, 0x8c, 0xbc, 0x7e, 0xd8, 0xa7, 0x27, 0xd9, 0xc2, 0x5d,
	0xf1, 0x93, 0x83, 0xb4, 0x5f, 0x18, 0xa5, 0x0e, 0x08, 0xc1, 0x1d, 0x05, 0x90, 0xda, 0xed, 0x71,
	0xdb, 0xeb, 0x87, 0xe6, 0x4c, 0xa6, 0x78, 0xa8, 0xc4, 0x56, 0x52, 0x88, 0xa8, 0x48, 0x21, 0xda,
	0xca, 0xdc, 0xee, 0x98, 0x9c, 0x41, 0xfd, 0x79, 0x21, 0x52, 0x58, 0xcb, 0x9c, 0xcc, 0xb0, 0x45,
	0x63, 0x2e, 0x38, 0xa3, 0x71, 0x1e, 0xb1, 0x84, 0x46, 0x34, 0x67, 0xc0, 0xb5, 0xdd, 0x35, 0x30,
	0x39, 0x69, 0xcb, 0x84, 0x2e, 0x8c, 0x42, 0x3e, 0xe2, 0x97, 0x3f, 0x40, 0x2a, 0x26, 0xb8, 0xdd,
	0x1b, 0x23, 0xef, 0xd5, 0xfc, 0x9d, 0x7f, 0x69, 0x83, 0xff, 0xb5, 0x46, 0x96, 0xfc, 0xbb, 0x08,
	0x1b, 0x7e, 0xe2, 0xe2, 0xd1, 0xb5, 0xa9, 0x43, 0x50, 0xa5, 0xe0, 0x0a, 0x26, 0x77, 0x08, 0x93,
	0x95, 0xca, 0xd6, 0x65, 0x1a, 0x6b, 0x30, 0x72, 0x95, 0xff, 0xcc, 0x4d, 0x19, 0x61, 0xe7, 0x72,
	0xe6, 0xc6, 0x92, 0xf9, 0x7f, 0x84, 0xdb, 0x2b, 0x95, 0x11, 0x81, 0xdf, 0x5c, 0xde, 0x16, 0xef,
	0x5a, 0x93, 0x6b, 0x0e, 0x3b, 0xb3, 0x5b, 0xc9, 0xa6, 0x31, 0x61, 0x78, 0x78, 0xfe, 0x1f, 0xa6,
	0x4f, 0x14, 0x39, 0xe3, 0x1c, 0xff, 0x36, 0xae, 0x69, 0xf5, 0xf9, 0xd3, 0x9f, 0x83, 0x8b, 0xf6,
	0x07, 0x17, 0xfd, 0x3b, 0xb8, 0xe8, 0xf7, 0xd1, 0x6d, 0xed, 0x8f, 0x6e, 0xeb, 0xef, 0xd1, 0x6d,
	0x7d, 0x7b, 0x9f, 0x31, 0xbd, 0xd9, 0x26, 0x3e, 0x15, 0x45, 0xf0, 0x58, 0xb3, 0x5a, 0xa4, 0x60,
	0x67, 0xde, 0xfa, 0x57, 0x09, 0x2a, 0xe9, 0x99, 0xad, 0xfa, 0xf0, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xb5, 0xa1, 0x50, 0xe3, 0x9e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RegisterChainName(ctx context.Context, in *MsgRegisterChainName, opts ...grpc.CallOption) (*MsgRegisterChainNameResponse, error)
	UpdateChainInfo(ctx context.Context, in *MsgUpdateChainInfo, opts ...grpc.CallOption) (*MsgUpdateChainInfoResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterChainName(ctx context.Context, in *MsgRegisterChainName, opts ...grpc.CallOption) (*MsgRegisterChainNameResponse, error) {
	out := new(MsgRegisterChainNameResponse)
	err := c.cc.Invoke(ctx, "/tendermint.cns.cns.Msg/RegisterChainName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateChainInfo(ctx context.Context, in *MsgUpdateChainInfo, opts ...grpc.CallOption) (*MsgUpdateChainInfoResponse, error) {
	out := new(MsgUpdateChainInfoResponse)
	err := c.cc.Invoke(ctx, "/tendermint.cns.cns.Msg/UpdateChainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RegisterChainName(context.Context, *MsgRegisterChainName) (*MsgRegisterChainNameResponse, error)
	UpdateChainInfo(context.Context, *MsgUpdateChainInfo) (*MsgUpdateChainInfoResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterChainName(ctx context.Context, req *MsgRegisterChainName) (*MsgRegisterChainNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterChainName not implemented")
}
func (*UnimplementedMsgServer) UpdateChainInfo(ctx context.Context, req *MsgUpdateChainInfo) (*MsgUpdateChainInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChainInfo not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterChainName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterChainName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterChainName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.cns.cns.Msg/RegisterChainName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterChainName(ctx, req.(*MsgRegisterChainName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateChainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateChainInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateChainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.cns.cns.Msg/UpdateChainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateChainInfo(ctx, req.(*MsgUpdateChainInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.cns.cns.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterChainName",
			Handler:    _Msg_RegisterChainName_Handler,
		},
		{
			MethodName: "UpdateChainInfo",
			Handler:    _Msg_UpdateChainInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cns/tx.proto",
}

func (m *MsgRegisterChainName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterChainName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterChainName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != nil {
		{
			size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.CanonicalIbcClient) > 0 {
		i -= len(m.CanonicalIbcClient)
		copy(dAtA[i:], m.CanonicalIbcClient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CanonicalIbcClient)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SourceCodeUrl) > 0 {
		i -= len(m.SourceCodeUrl)
		copy(dAtA[i:], m.SourceCodeUrl)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceCodeUrl)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Seed) > 0 {
		for iNdEx := len(m.Seed) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Seed[iNdEx])
			copy(dAtA[i:], m.Seed[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Seed[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterChainNameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterChainNameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterChainNameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != nil {
		{
			size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.CanonicalIbcClient) > 0 {
		i -= len(m.CanonicalIbcClient)
		copy(dAtA[i:], m.CanonicalIbcClient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CanonicalIbcClient)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SourceCodeUrl) > 0 {
		i -= len(m.SourceCodeUrl)
		copy(dAtA[i:], m.SourceCodeUrl)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceCodeUrl)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Seed) > 0 {
		for iNdEx := len(m.Seed) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Seed[iNdEx])
			copy(dAtA[i:], m.Seed[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Seed[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateChainInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateChainInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateChainInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterChainName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Seed) > 0 {
		for _, s := range m.Seed {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.SourceCodeUrl)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CanonicalIbcClient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Version != nil {
		l = m.Version.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterChainNameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Seed) > 0 {
		for _, s := range m.Seed {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.SourceCodeUrl)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CanonicalIbcClient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Version != nil {
		l = m.Version.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateChainInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterChainName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterChainName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterChainName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = append(m.Seed, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceCodeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceCodeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalIbcClient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CanonicalIbcClient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Version == nil {
				m.Version = &VersionInfo{}
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRegisterChainNameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterChainNameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterChainNameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = append(m.Seed, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceCodeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceCodeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalIbcClient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CanonicalIbcClient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Version == nil {
				m.Version = &VersionInfo{}
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateChainInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateChainInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateChainInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
