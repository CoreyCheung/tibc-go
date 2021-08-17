// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tibc/core/routing/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryRoutingRulesRequest is the request type for the
// RoutingRules RPC method
type QueryRoutingRulesRequest struct {
}

func (m *QueryRoutingRulesRequest) Reset()         { *m = QueryRoutingRulesRequest{} }
func (m *QueryRoutingRulesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRoutingRulesRequest) ProtoMessage()    {}
func (*QueryRoutingRulesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af620a084c71311, []int{0}
}
func (m *QueryRoutingRulesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRoutingRulesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRoutingRulesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRoutingRulesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRoutingRulesRequest.Merge(m, src)
}
func (m *QueryRoutingRulesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRoutingRulesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRoutingRulesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRoutingRulesRequest proto.InternalMessageInfo

// QueryRoutingRulesResponse defines the routing rules query response
type QueryRoutingRulesResponse struct {
	// rule string array
	Rules []string `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (m *QueryRoutingRulesResponse) Reset()         { *m = QueryRoutingRulesResponse{} }
func (m *QueryRoutingRulesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRoutingRulesResponse) ProtoMessage()    {}
func (*QueryRoutingRulesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af620a084c71311, []int{1}
}
func (m *QueryRoutingRulesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRoutingRulesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRoutingRulesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRoutingRulesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRoutingRulesResponse.Merge(m, src)
}
func (m *QueryRoutingRulesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRoutingRulesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRoutingRulesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRoutingRulesResponse proto.InternalMessageInfo

func (m *QueryRoutingRulesResponse) GetRules() []string {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRoutingRulesRequest)(nil), "tibc.core.routing.v1.QueryRoutingRulesRequest")
	proto.RegisterType((*QueryRoutingRulesResponse)(nil), "tibc.core.routing.v1.QueryRoutingRulesResponse")
}

func init() { proto.RegisterFile("tibc/core/routing/v1/query.proto", fileDescriptor_7af620a084c71311) }

var fileDescriptor_7af620a084c71311 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xc9, 0x4c, 0x4a,
	0xd6, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0xca, 0x2f, 0x2d, 0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x33,
	0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x01, 0xa9,
	0xd0, 0x03, 0xa9, 0xd0, 0x83, 0xaa, 0xd0, 0x2b, 0x33, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf,
	0x2b, 0x86, 0xe8, 0x51, 0x92, 0xe2, 0x92, 0x08, 0x04, 0x19, 0x11, 0x04, 0xd1, 0x10, 0x54, 0x9a,
	0x93, 0x5a, 0x1c, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xa2, 0x64, 0xc8, 0x25, 0x89, 0x45, 0xae,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x84, 0x8b, 0xb5, 0x08, 0x24, 0x20, 0xc1, 0xa8, 0xc0,
	0xac, 0xc1, 0x19, 0x04, 0xe1, 0x18, 0xad, 0x66, 0xe4, 0x62, 0x05, 0xeb, 0x11, 0x5a, 0xc8, 0xc8,
	0xc5, 0x83, 0xac, 0x51, 0x48, 0x4f, 0x0f, 0x9b, 0xf3, 0xf4, 0x70, 0xd9, 0x2e, 0xa5, 0x4f, 0xb4,
	0x7a, 0x88, 0x8b, 0x94, 0x0c, 0x9a, 0x2e, 0x3f, 0x99, 0xcc, 0xa4, 0x25, 0xa4, 0xa1, 0x8f, 0x2d,
	0xa4, 0x92, 0x52, 0x4b, 0x12, 0x0d, 0x61, 0xfc, 0x78, 0xb0, 0x6b, 0x9d, 0x22, 0x4f, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x3e, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f,
	0x39, 0x3f, 0x57, 0x3f, 0x29, 0x33, 0x31, 0x2f, 0x2b, 0x33, 0x35, 0x31, 0x13, 0x6c, 0xae, 0x6e,
	0x7a, 0xbe, 0x7e, 0x6e, 0x7e, 0x0a, 0x48, 0x3f, 0x92, 0x3d, 0x46, 0x66, 0xba, 0x30, 0xab, 0x4a,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xc1, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x11,
	0x35, 0x26, 0xf8, 0xb6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// RoutingRules queries all routing rules.
	RoutingRules(ctx context.Context, in *QueryRoutingRulesRequest, opts ...grpc.CallOption) (*QueryRoutingRulesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) RoutingRules(ctx context.Context, in *QueryRoutingRulesRequest, opts ...grpc.CallOption) (*QueryRoutingRulesResponse, error) {
	out := new(QueryRoutingRulesResponse)
	err := c.cc.Invoke(ctx, "/tibc.core.routing.v1.Query/RoutingRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// RoutingRules queries all routing rules.
	RoutingRules(context.Context, *QueryRoutingRulesRequest) (*QueryRoutingRulesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) RoutingRules(ctx context.Context, req *QueryRoutingRulesRequest) (*QueryRoutingRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoutingRules not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_RoutingRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRoutingRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RoutingRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tibc.core.routing.v1.Query/RoutingRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RoutingRules(ctx, req.(*QueryRoutingRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tibc.core.routing.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RoutingRules",
			Handler:    _Query_RoutingRules_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tibc/core/routing/v1/query.proto",
}

func (m *QueryRoutingRulesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRoutingRulesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRoutingRulesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryRoutingRulesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRoutingRulesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRoutingRulesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for iNdEx := len(m.Rules) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Rules[iNdEx])
			copy(dAtA[i:], m.Rules[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Rules[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRoutingRulesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryRoutingRulesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, s := range m.Rules {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRoutingRulesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRoutingRulesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRoutingRulesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRoutingRulesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRoutingRulesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRoutingRulesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)