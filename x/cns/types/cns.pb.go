// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/cns/cns.proto

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
	return fileDescriptor_b17ae4027e1155ce, []int{0}
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
	return fileDescriptor_b17ae4027e1155ce, []int{1}
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

func init() { proto.RegisterFile("proto/cns/cns.proto", fileDescriptor_b17ae4027e1155ce) }

var fileDescriptor_b17ae4027e1155ce = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xbd, 0x8e, 0xd4, 0x30,
	0x14, 0x85, 0xc7, 0x9b, 0xd9, 0xdd, 0x89, 0x07, 0x04, 0xba, 0x4c, 0x61, 0xad, 0x44, 0x26, 0xac,
	0x04, 0x4a, 0x81, 0xb2, 0x08, 0x2a, 0xb6, 0x41, 0x62, 0x1a, 0xb6, 0xa1, 0x88, 0x04, 0x05, 0x4d,
	0xe4, 0x38, 0x97, 0xc4, 0x52, 0x62, 0x8f, 0x6c, 0x0f, 0x2c, 0x0d, 0xcf, 0x40, 0x49, 0xb9, 0x6f,
	0xc0, 0x6b, 0x50, 0x6e, 0x49, 0x89, 0x66, 0x1a, 0x1e, 0x03, 0xd9, 0x61, 0x7e, 0x24, 0x8a, 0x44,
	0x3e, 0xe7, 0x7e, 0x3a, 0xb9, 0x39, 0x09, 0x7d, 0xb0, 0x34, 0xda, 0xe9, 0x0b, 0xa1, 0xac, 0xbf,
	0xf2, 0xa0, 0x00, 0x1c, 0xaa, 0x1a, 0x4d, 0x2f, 0x95, 0xcb, 0xbd, 0x2b, 0x94, 0x3d, 0x9b, 0x35,
	0xba, 0xd1, 0x03, 0xec, 0x4f, 0x03, 0x79, 0xfe, 0xe3, 0x88, 0xc6, 0x8b, 0x96, 0x4b, 0x75, 0xa5,
	0x3e, 0x6a, 0x78, 0x48, 0xa9, 0xf0, 0xa2, 0x54, 0xbc, 0x47, 0x46, 0x52, 0x92, 0xc5, 0x45, 0x1c,
	0x9c, 0xb7, 0xbc, 0x47, 0x48, 0x28, 0xc5, 0xeb, 0xa5, 0x34, 0xdc, 0x49, 0xad, 0xd8, 0x51, 0x4a,
	0xb2, 0xa8, 0x38, 0x70, 0xe0, 0x8c, 0x4e, 0x7a, 0x74, 0xbc, 0xe6, 0x8e, 0xb3, 0x28, 0x8d, 0xb2,
	0xb8, 0xd8, 0x69, 0x98, 0xd1, 0x63, 0xfd, 0x59, 0xa1, 0x61, 0xe3, 0x90, 0x3a, 0x08, 0x78, 0x46,
	0x67, 0x82, 0x2b, 0xad, 0xa4, 0xe0, 0x5d, 0x29, 0x2b, 0x51, 0x8a, 0x4e, 0xa2, 0x72, 0xec, 0x38,
	0x40, 0xb0, 0x9b, 0x5d, 0x55, 0x62, 0x11, 0x26, 0x00, 0x74, 0x6c, 0x11, 0x6b, 0x76, 0x12, 0xf2,
	0xc3, 0x19, 0x9e, 0xd0, 0x7b, 0x56, 0xaf, 0x8c, 0xc0, 0x52, 0xe8, 0x1a, 0xcb, 0x95, 0xe9, 0xd8,
	0x69, 0x08, 0xb8, 0x3b, 0xd8, 0x0b, 0x5d, 0xe3, 0x3b, 0xd3, 0xc1, 0x4b, 0x7a, 0xfa, 0x09, 0x8d,
	0xf5, 0xcb, 0x4f, 0x52, 0x92, 0x4d, 0x9f, 0xcf, 0xf3, 0xff, 0x8b, 0xca, 0xdf, 0x0f, 0x88, 0x2f,
	0xa4, 0xd8, 0xf2, 0x97, 0x93, 0xef, 0x37, 0x73, 0xf2, 0xe7, 0x66, 0x4e, 0xce, 0xbf, 0xd2, 0xe9,
	0x01, 0x01, 0x6c, 0x9f, 0x49, 0x42, 0x21, 0x5b, 0x09, 0x4f, 0x29, 0xec, 0xb6, 0xea, 0x7b, 0xe9,
	0xca, 0x96, 0xdb, 0x36, 0xb4, 0x16, 0x17, 0xf7, 0xb7, 0x8b, 0xf9, 0xc1, 0x1b, 0x6e, 0x5b, 0x78,
	0x44, 0xef, 0x34, 0xa8, 0xd0, 0x4a, 0x3b, 0x70, 0x51, 0xe0, 0xa6, 0xff, 0x3c, 0x8f, 0x5c, 0x8e,
	0xfd, 0xf3, 0x5f, 0xbf, 0xfa, 0xb9, 0x4e, 0xc8, 0xed, 0x3a, 0x21, 0xbf, 0xd7, 0x09, 0xf9, 0xb6,
	0x49, 0x46, 0xb7, 0x9b, 0x64, 0xf4, 0x6b, 0x93, 0x8c, 0x3e, 0x3c, 0x6e, 0xa4, 0x6b, 0x57, 0x55,
	0x2e, 0x74, 0x7f, 0xb1, 0x7f, 0xaf, 0xf0, 0x6b, 0x5c, 0x87, 0xbb, 0xfb, 0xb2, 0x44, 0x5b, 0x9d,
	0x84, 0x2f, 0xff, 0xe2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xe5, 0x3c, 0xe0, 0x3a, 0x02,
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
