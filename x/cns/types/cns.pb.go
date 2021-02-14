// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cns/cns.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

//TODO(sahith): Add json and yaml flags
type ChainInfo struct {
	ChainName          string       `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	Expiration         int64        `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Metadata           []string     `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"`
	Owner              string       `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	CanonicalIbcClient string       `protobuf:"bytes,5,opt,name=canonical_ibc_client,json=canonicalIbcClient,proto3" json:"canonical_ibc_client,omitempty"`
	Seed               []string     `protobuf:"bytes,6,rep,name=seed,proto3" json:"seed,omitempty"`
	SourceCodeUrl      string       `protobuf:"bytes,7,opt,name=source_code_url,json=sourceCodeUrl,proto3" json:"source_code_url,omitempty"`
	Version            *VersionInfo `protobuf:"bytes,8,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ChainInfo) Reset()         { *m = ChainInfo{} }
func (m *ChainInfo) String() string { return proto.CompactTextString(m) }
func (*ChainInfo) ProtoMessage()    {}
func (*ChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4614c872fdf07012, []int{0}
}
func (m *ChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainInfo.Merge(m, src)
}
func (m *ChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *ChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChainInfo proto.InternalMessageInfo

func (m *ChainInfo) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *ChainInfo) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *ChainInfo) GetMetadata() []string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ChainInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ChainInfo) GetCanonicalIbcClient() string {
	if m != nil {
		return m.CanonicalIbcClient
	}
	return ""
}

func (m *ChainInfo) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *ChainInfo) GetSourceCodeUrl() string {
	if m != nil {
		return m.SourceCodeUrl
	}
	return ""
}

func (m *ChainInfo) GetVersion() *VersionInfo {
	if m != nil {
		return m.Version
	}
	return nil
}

type VersionInfo struct {
	Version          int64  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	SourceCommitHash string `protobuf:"bytes,2,opt,name=source_commit_hash,json=sourceCommitHash,proto3" json:"source_commit_hash,omitempty"`
	GenesisHash      string `protobuf:"bytes,3,opt,name=genesis_hash,json=genesisHash,proto3" json:"genesis_hash,omitempty"`
}

func (m *VersionInfo) Reset()         { *m = VersionInfo{} }
func (m *VersionInfo) String() string { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()    {}
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4614c872fdf07012, []int{1}
}
func (m *VersionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionInfo.Merge(m, src)
}
func (m *VersionInfo) XXX_Size() int {
	return m.Size()
}
func (m *VersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VersionInfo proto.InternalMessageInfo

func (m *VersionInfo) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VersionInfo) GetSourceCommitHash() string {
	if m != nil {
		return m.SourceCommitHash
	}
	return ""
}

func (m *VersionInfo) GetGenesisHash() string {
	if m != nil {
		return m.GenesisHash
	}
	return ""
}

func init() {
	proto.RegisterType((*ChainInfo)(nil), "tendermint.cns.cns.ChainInfo")
	proto.RegisterType((*VersionInfo)(nil), "tendermint.cns.cns.VersionInfo")
}

func init() { proto.RegisterFile("cns/cns.proto", fileDescriptor_4614c872fdf07012) }

var fileDescriptor_4614c872fdf07012 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x18, 0x86, 0xeb, 0x6b, 0xef, 0xae, 0x71, 0x39, 0x81, 0x3e, 0x75, 0xb0, 0x4e, 0x22, 0x0d, 0x27,
	0x81, 0x32, 0xa0, 0x04, 0xc1, 0xc4, 0x2d, 0x48, 0x74, 0xe1, 0x16, 0x86, 0x48, 0x30, 0xb0, 0x44,
	0x8e, 0xf3, 0x91, 0x58, 0x4a, 0xec, 0xca, 0x76, 0xe1, 0x58, 0xf8, 0x0d, 0x8c, 0x8c, 0xf7, 0x0f,
	0xf8, 0x1b, 0x8c, 0x37, 0x32, 0xa2, 0x76, 0xe1, 0x67, 0x20, 0x3b, 0xb4, 0x57, 0x89, 0x25, 0xf2,
	0xf7, 0xbe, 0x8f, 0xde, 0x7c, 0x79, 0x63, 0x7a, 0x26, 0x94, 0xcd, 0x85, 0xb2, 0xd9, 0xca, 0x68,
	0xa7, 0x01, 0x1c, 0xaa, 0x1a, 0x4d, 0x2f, 0x95, 0xcb, 0xbc, 0x2a, 0x94, 0x3d, 0x9f, 0x37, 0xba,
	0xd1, 0xc1, 0xce, 0xfd, 0x69, 0x20, 0x2f, 0x7e, 0x1c, 0xd1, 0x68, 0xd9, 0x72, 0xa9, 0xae, 0xd4,
	0x47, 0x0d, 0x0f, 0x29, 0x15, 0x7e, 0x28, 0x15, 0xef, 0x91, 0x91, 0x84, 0xa4, 0x51, 0x11, 0x05,
	0xe5, 0x2d, 0xef, 0x11, 0x62, 0x4a, 0xf1, 0x7a, 0x25, 0x0d, 0x77, 0x52, 0x2b, 0x76, 0x94, 0x90,
	0x74, 0x5c, 0x1c, 0x28, 0x70, 0x4e, 0xa7, 0x3d, 0x3a, 0x5e, 0x73, 0xc7, 0xd9, 0x38, 0x19, 0xa7,
	0x51, 0xb1, 0x9f, 0x61, 0x4e, 0x8f, 0xf5, 0x67, 0x85, 0x86, 0x4d, 0x42, 0xea, 0x30, 0xc0, 0x33,
	0x3a, 0x17, 0x5c, 0x69, 0x25, 0x05, 0xef, 0x4a, 0x59, 0x89, 0x52, 0x74, 0x12, 0x95, 0x63, 0xc7,
	0x01, 0x82, 0xbd, 0x77, 0x55, 0x89, 0x65, 0x70, 0x00, 0xe8, 0xc4, 0x22, 0xd6, 0xec, 0x24, 0xe4,
	0x87, 0x33, 0x3c, 0xa1, 0xf7, 0xad, 0x5e, 0x1b, 0x81, 0xa5, 0xd0, 0x35, 0x96, 0x6b, 0xd3, 0xb1,
	0xd3, 0x10, 0x70, 0x36, 0xc8, 0x4b, 0x5d, 0xe3, 0x3b, 0xd3, 0xc1, 0x4b, 0x7a, 0xfa, 0x09, 0x8d,
	0xf5, 0xcb, 0x4f, 0x13, 0x92, 0xce, 0x9e, 0x2f, 0xb2, 0xff, 0x8b, 0xca, 0xde, 0x0f, 0x88, 0x2f,
	0xa4, 0xd8, 0xf1, 0x97, 0xd3, 0xef, 0x37, 0x0b, 0xf2, 0xe7, 0x66, 0x41, 0x2e, 0xbe, 0xd2, 0xd9,
	0x01, 0x01, 0xec, 0x2e, 0x93, 0x84, 0x42, 0x76, 0x23, 0x3c, 0xa5, 0xb0, 0xdf, 0xaa, 0xef, 0xa5,
	0x2b, 0x5b, 0x6e, 0xdb, 0xd0, 0x5a, 0x54, 0x3c, 0xd8, 0x2d, 0xe6, 0x8d, 0x37, 0xdc, 0xb6, 0xf0,
	0x88, 0xde, 0x6b, 0x50, 0xa1, 0x95, 0x76, 0xe0, 0xc6, 0x81, 0x9b, 0xfd, 0xd3, 0x3c, 0x72, 0x39,
	0xf1, 0xef, 0x7f, 0xfd, 0xea, 0xe7, 0x26, 0x26, 0xb7, 0x9b, 0x98, 0xfc, 0xde, 0xc4, 0xe4, 0xdb,
	0x36, 0x1e, 0xdd, 0x6e, 0xe3, 0xd1, 0xaf, 0x6d, 0x3c, 0xfa, 0xf0, 0xb8, 0x91, 0xae, 0x5d, 0x57,
	0x99, 0xd0, 0x7d, 0x7e, 0xf7, 0x5d, 0xfe, 0x5a, 0xe4, 0xd7, 0xe1, 0xe9, 0xbe, 0xac, 0xd0, 0x56,
	0x27, 0xe1, 0xcf, 0xbf, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x87, 0xc4, 0x01, 0x97, 0x34, 0x02,
	0x00, 0x00,
}

func (this *ChainInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChainInfo)
	if !ok {
		that2, ok := that.(ChainInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ChainName != that1.ChainName {
		return false
	}
	if this.Expiration != that1.Expiration {
		return false
	}
	if len(this.Metadata) != len(that1.Metadata) {
		return false
	}
	for i := range this.Metadata {
		if this.Metadata[i] != that1.Metadata[i] {
			return false
		}
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.CanonicalIbcClient != that1.CanonicalIbcClient {
		return false
	}
	if len(this.Seed) != len(that1.Seed) {
		return false
	}
	for i := range this.Seed {
		if this.Seed[i] != that1.Seed[i] {
			return false
		}
	}
	if this.SourceCodeUrl != that1.SourceCodeUrl {
		return false
	}
	if !this.Version.Equal(that1.Version) {
		return false
	}
	return true
}
func (this *VersionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VersionInfo)
	if !ok {
		that2, ok := that.(VersionInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.SourceCommitHash != that1.SourceCommitHash {
		return false
	}
	if this.GenesisHash != that1.GenesisHash {
		return false
	}
	return true
}
func (m *ChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintCns(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.SourceCodeUrl) > 0 {
		i -= len(m.SourceCodeUrl)
		copy(dAtA[i:], m.SourceCodeUrl)
		i = encodeVarintCns(dAtA, i, uint64(len(m.SourceCodeUrl)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Seed) > 0 {
		for iNdEx := len(m.Seed) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Seed[iNdEx])
			copy(dAtA[i:], m.Seed[iNdEx])
			i = encodeVarintCns(dAtA, i, uint64(len(m.Seed[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.CanonicalIbcClient) > 0 {
		i -= len(m.CanonicalIbcClient)
		copy(dAtA[i:], m.CanonicalIbcClient)
		i = encodeVarintCns(dAtA, i, uint64(len(m.CanonicalIbcClient)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCns(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Metadata) > 0 {
		for iNdEx := len(m.Metadata) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Metadata[iNdEx])
			copy(dAtA[i:], m.Metadata[iNdEx])
			i = encodeVarintCns(dAtA, i, uint64(len(m.Metadata[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Expiration != 0 {
		i = encodeVarintCns(dAtA, i, uint64(m.Expiration))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintCns(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VersionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VersionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GenesisHash) > 0 {
		i -= len(m.GenesisHash)
		copy(dAtA[i:], m.GenesisHash)
		i = encodeVarintCns(dAtA, i, uint64(len(m.GenesisHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceCommitHash) > 0 {
		i -= len(m.SourceCommitHash)
		copy(dAtA[i:], m.SourceCommitHash)
		i = encodeVarintCns(dAtA, i, uint64(len(m.SourceCommitHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarintCns(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCns(dAtA []byte, offset int, v uint64) int {
	offset -= sovCns(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovCns(uint64(l))
	}
	if m.Expiration != 0 {
		n += 1 + sovCns(uint64(m.Expiration))
	}
	if len(m.Metadata) > 0 {
		for _, s := range m.Metadata {
			l = len(s)
			n += 1 + l + sovCns(uint64(l))
		}
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCns(uint64(l))
	}
	l = len(m.CanonicalIbcClient)
	if l > 0 {
		n += 1 + l + sovCns(uint64(l))
	}
	if len(m.Seed) > 0 {
		for _, s := range m.Seed {
			l = len(s)
			n += 1 + l + sovCns(uint64(l))
		}
	}
	l = len(m.SourceCodeUrl)
	if l > 0 {
		n += 1 + l + sovCns(uint64(l))
	}
	if m.Version != nil {
		l = m.Version.Size()
		n += 1 + l + sovCns(uint64(l))
	}
	return n
}

func (m *VersionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovCns(uint64(m.Version))
	}
	l = len(m.SourceCommitHash)
	if l > 0 {
		n += 1 + l + sovCns(uint64(l))
	}
	l = len(m.GenesisHash)
	if l > 0 {
		n += 1 + l + sovCns(uint64(l))
	}
	return n
}

func sovCns(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCns(x uint64) (n int) {
	return sovCns(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCns
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
			return fmt.Errorf("proto: ChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
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
				return ErrInvalidLengthCns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			m.Expiration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
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
				return ErrInvalidLengthCns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = append(m.Metadata, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
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
				return ErrInvalidLengthCns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalIbcClient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
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
				return ErrInvalidLengthCns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CanonicalIbcClient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
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
				return ErrInvalidLengthCns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = append(m.Seed, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceCodeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
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
				return ErrInvalidLengthCns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceCodeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
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
				return ErrInvalidLengthCns
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCns
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
			skippy, err := skipCns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCns
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
func (m *VersionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCns
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
			return fmt.Errorf("proto: VersionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceCommitHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
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
				return ErrInvalidLengthCns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceCommitHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCns
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
				return ErrInvalidLengthCns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCns
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
func skipCns(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCns
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
					return 0, ErrIntOverflowCns
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
					return 0, ErrIntOverflowCns
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
				return 0, ErrInvalidLengthCns
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCns
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCns
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCns        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCns          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCns = fmt.Errorf("proto: unexpected end of group")
)