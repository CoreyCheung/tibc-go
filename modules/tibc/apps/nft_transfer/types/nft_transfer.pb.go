// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tibc/apps/nft_transfer/v1/nft_transfer.proto

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

type NonFungibleTokenPacketData struct {
	// the class to which the NFT to be transferred belongs
	Class string `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	// the nft id
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// the address defined by NFT outside the chain
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	// the nft sender
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
	// the nft receiver
	Receiver string `protobuf:"bytes,5,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// identify whether it is far away from the source chain
	AwayFromOrigin bool `protobuf:"varint,6,opt,name=away_from_origin,json=awayFromOrigin,proto3" json:"away_from_origin,omitempty"`
}

func (m *NonFungibleTokenPacketData) Reset()         { *m = NonFungibleTokenPacketData{} }
func (m *NonFungibleTokenPacketData) String() string { return proto.CompactTextString(m) }
func (*NonFungibleTokenPacketData) ProtoMessage()    {}
func (*NonFungibleTokenPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_323a0d65d2518331, []int{0}
}
func (m *NonFungibleTokenPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NonFungibleTokenPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NonFungibleTokenPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NonFungibleTokenPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonFungibleTokenPacketData.Merge(m, src)
}
func (m *NonFungibleTokenPacketData) XXX_Size() int {
	return m.Size()
}
func (m *NonFungibleTokenPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonFungibleTokenPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_NonFungibleTokenPacketData proto.InternalMessageInfo

func (m *NonFungibleTokenPacketData) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *NonFungibleTokenPacketData) GetAwayFromOrigin() bool {
	if m != nil {
		return m.AwayFromOrigin
	}
	return false
}

// ClassTrace contains the base class for TICS30 Non fungible tokens and the
// source tracing information path.
type ClassTrace struct {
	// path defines the chain of sourceChain/destChain
	// identifiers used for tracing the source of the Non fungible token.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// base class of the relayed non fungible token.
	BaseClass string `protobuf:"bytes,2,opt,name=base_class,json=baseClass,proto3" json:"base_class,omitempty"`
}

func (m *ClassTrace) Reset()         { *m = ClassTrace{} }
func (m *ClassTrace) String() string { return proto.CompactTextString(m) }
func (*ClassTrace) ProtoMessage()    {}
func (*ClassTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_323a0d65d2518331, []int{1}
}
func (m *ClassTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassTrace.Merge(m, src)
}
func (m *ClassTrace) XXX_Size() int {
	return m.Size()
}
func (m *ClassTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassTrace.DiscardUnknown(m)
}

var xxx_messageInfo_ClassTrace proto.InternalMessageInfo

func (m *ClassTrace) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ClassTrace) GetBaseClass() string {
	if m != nil {
		return m.BaseClass
	}
	return ""
}

func init() {
	proto.RegisterType((*NonFungibleTokenPacketData)(nil), "tibc.apps.nft_transfer.v1.NonFungibleTokenPacketData")
	proto.RegisterType((*ClassTrace)(nil), "tibc.apps.nft_transfer.v1.ClassTrace")
}

func init() {
	proto.RegisterFile("tibc/apps/nft_transfer/v1/nft_transfer.proto", fileDescriptor_323a0d65d2518331)
}

var fileDescriptor_323a0d65d2518331 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x6a, 0xc2, 0x40,
	0x14, 0x85, 0x1d, 0xff, 0xd0, 0x59, 0x88, 0x0c, 0x52, 0x52, 0xa1, 0x41, 0x5c, 0xb9, 0x68, 0x33,
	0x48, 0x1f, 0xa0, 0xf4, 0x07, 0x97, 0x6d, 0x11, 0x57, 0xed, 0x22, 0x4c, 0x92, 0x6b, 0x9c, 0x6a,
	0x66, 0xc2, 0xcc, 0xc4, 0xe2, 0x5b, 0xf4, 0x59, 0xfa, 0x14, 0x5d, 0xba, 0xec, 0xb2, 0x98, 0x17,
	0x29, 0x33, 0x91, 0x82, 0x8b, 0xee, 0xce, 0xf9, 0xee, 0xb9, 0x70, 0xb8, 0x17, 0x5f, 0x1a, 0x1e,
	0xc5, 0x94, 0xe5, 0xb9, 0xa6, 0x62, 0x69, 0x42, 0xa3, 0x98, 0xd0, 0x4b, 0x50, 0x74, 0x3b, 0x3d,
	0xf1, 0x41, 0xae, 0xa4, 0x91, 0xe4, 0xdc, 0xa6, 0x03, 0x9b, 0x0e, 0x4e, 0xa6, 0xdb, 0xe9, 0x70,
	0x90, 0xca, 0x54, 0xba, 0x14, 0xb5, 0xaa, 0x5a, 0x18, 0x7f, 0x22, 0x3c, 0x7c, 0x94, 0x62, 0x56,
	0x88, 0x94, 0x47, 0x1b, 0x58, 0xc8, 0x35, 0x88, 0x67, 0x16, 0xaf, 0xc1, 0x3c, 0x30, 0xc3, 0xc8,
	0x00, 0xb7, 0xe2, 0x0d, 0xd3, 0xda, 0x43, 0x23, 0x34, 0xe9, 0xce, 0x2b, 0x43, 0x7a, 0xb8, 0xce,
	0x13, 0xaf, 0xee, 0x50, 0x9d, 0x27, 0xa4, 0x8f, 0x1b, 0x85, 0xe2, 0x5e, 0xc3, 0x01, 0x2b, 0xc9,
	0x19, 0x6e, 0x6b, 0x10, 0x09, 0x28, 0xaf, 0xe9, 0xe0, 0xd1, 0x91, 0x21, 0xee, 0x28, 0x88, 0x81,
	0x6f, 0x41, 0x79, 0x2d, 0x37, 0xf9, 0xf3, 0x64, 0x82, 0xfb, 0xec, 0x9d, 0xed, 0xc2, 0xa5, 0x92,
	0x59, 0x28, 0x15, 0x4f, 0xb9, 0xf0, 0xda, 0x23, 0x34, 0xe9, 0xcc, 0x7b, 0x96, 0xcf, 0x94, 0xcc,
	0x9e, 0x1c, 0x1d, 0xdf, 0x60, 0x7c, 0x6f, 0x8b, 0x2c, 0x14, 0x8b, 0x81, 0x10, 0xdc, 0xcc, 0x99,
	0x59, 0x1d, 0x2b, 0x3a, 0x4d, 0x2e, 0x30, 0x8e, 0x98, 0x86, 0xb0, 0x2a, 0x5f, 0x35, 0xed, 0x5a,
	0xe2, 0xf6, 0xee, 0x5e, 0xbf, 0x0e, 0x3e, 0xda, 0x1f, 0x7c, 0xf4, 0x73, 0xf0, 0xd1, 0x47, 0xe9,
	0xd7, 0xf6, 0xa5, 0x5f, 0xfb, 0x2e, 0xfd, 0xda, 0xcb, 0x6d, 0xca, 0xcd, 0xaa, 0x88, 0x82, 0x58,
	0x66, 0x34, 0xe2, 0x4c, 0xbc, 0x71, 0x60, 0x9c, 0xda, 0xab, 0x5e, 0xa5, 0x92, 0x66, 0x32, 0x29,
	0x36, 0xa0, 0xe9, 0x3f, 0x3f, 0x31, 0xbb, 0x1c, 0x74, 0xd4, 0x76, 0x97, 0xbd, 0xfe, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x72, 0x28, 0xd7, 0xf3, 0xba, 0x01, 0x00, 0x00,
}

func (m *NonFungibleTokenPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NonFungibleTokenPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NonFungibleTokenPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AwayFromOrigin {
		i--
		if m.AwayFromOrigin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Class) > 0 {
		i -= len(m.Class)
		copy(dAtA[i:], m.Class)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Class)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClassTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BaseClass) > 0 {
		i -= len(m.BaseClass)
		copy(dAtA[i:], m.BaseClass)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.BaseClass)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintNftTransfer(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NonFungibleTokenPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Class)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	if m.AwayFromOrigin {
		n += 2
	}
	return n
}

func (m *ClassTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	l = len(m.BaseClass)
	if l > 0 {
		n += 1 + l + sovNftTransfer(uint64(l))
	}
	return n
}

func sovNftTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftTransfer(x uint64) (n int) {
	return sovNftTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NonFungibleTokenPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftTransfer
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
			return fmt.Errorf("proto: NonFungibleTokenPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NonFungibleTokenPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Class", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Class = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AwayFromOrigin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AwayFromOrigin = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNftTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftTransfer
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
func (m *ClassTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftTransfer
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
			return fmt.Errorf("proto: ClassTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseClass", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftTransfer
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
				return ErrInvalidLengthNftTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseClass = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftTransfer
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
func skipNftTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftTransfer
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
					return 0, ErrIntOverflowNftTransfer
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
					return 0, ErrIntOverflowNftTransfer
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
				return 0, ErrInvalidLengthNftTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftTransfer = fmt.Errorf("proto: unexpected end of group")
)
