// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identity/v1beta/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

type QueryUsernameFromAddressRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryUsernameFromAddressRequest) Reset()         { *m = QueryUsernameFromAddressRequest{} }
func (m *QueryUsernameFromAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUsernameFromAddressRequest) ProtoMessage()    {}
func (*QueryUsernameFromAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c88d5f879027c4, []int{0}
}
func (m *QueryUsernameFromAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUsernameFromAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUsernameFromAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUsernameFromAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUsernameFromAddressRequest.Merge(m, src)
}
func (m *QueryUsernameFromAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUsernameFromAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUsernameFromAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUsernameFromAddressRequest proto.InternalMessageInfo

func (m *QueryUsernameFromAddressRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryUsernameFromAddressResponse struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (m *QueryUsernameFromAddressResponse) Reset()         { *m = QueryUsernameFromAddressResponse{} }
func (m *QueryUsernameFromAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUsernameFromAddressResponse) ProtoMessage()    {}
func (*QueryUsernameFromAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c88d5f879027c4, []int{1}
}
func (m *QueryUsernameFromAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUsernameFromAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUsernameFromAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUsernameFromAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUsernameFromAddressResponse.Merge(m, src)
}
func (m *QueryUsernameFromAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUsernameFromAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUsernameFromAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUsernameFromAddressResponse proto.InternalMessageInfo

func (m *QueryUsernameFromAddressResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type QueryUsernameRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (m *QueryUsernameRequest) Reset()         { *m = QueryUsernameRequest{} }
func (m *QueryUsernameRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUsernameRequest) ProtoMessage()    {}
func (*QueryUsernameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c88d5f879027c4, []int{2}
}
func (m *QueryUsernameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUsernameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUsernameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUsernameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUsernameRequest.Merge(m, src)
}
func (m *QueryUsernameRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUsernameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUsernameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUsernameRequest proto.InternalMessageInfo

func (m *QueryUsernameRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type QueryUsernameResponse struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (m *QueryUsernameResponse) Reset()         { *m = QueryUsernameResponse{} }
func (m *QueryUsernameResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUsernameResponse) ProtoMessage()    {}
func (*QueryUsernameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80c88d5f879027c4, []int{3}
}
func (m *QueryUsernameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUsernameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUsernameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUsernameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUsernameResponse.Merge(m, src)
}
func (m *QueryUsernameResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUsernameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUsernameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUsernameResponse proto.InternalMessageInfo

func (m *QueryUsernameResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryUsernameFromAddressRequest)(nil), "spn.identity.v1beta1.QueryUsernameFromAddressRequest")
	proto.RegisterType((*QueryUsernameFromAddressResponse)(nil), "spn.identity.v1beta1.QueryUsernameFromAddressResponse")
	proto.RegisterType((*QueryUsernameRequest)(nil), "spn.identity.v1beta1.QueryUsernameRequest")
	proto.RegisterType((*QueryUsernameResponse)(nil), "spn.identity.v1beta1.QueryUsernameResponse")
}

func init() { proto.RegisterFile("identity/v1beta/querier.proto", fileDescriptor_80c88d5f879027c4) }

var fileDescriptor_80c88d5f879027c4 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0xb3, 0x85, 0x1a, 0xa7, 0x3c, 0x23, 0x84, 0x80, 0x6b, 0xb8, 0x4a, 0x22, 0xdc, 0x12,
	0x83, 0x69, 0x04, 0x41, 0x11, 0x7b, 0x03, 0x36, 0x76, 0x7b, 0xb9, 0x31, 0x6e, 0x71, 0xbb, 0x9b,
	0x9d, 0x3d, 0x31, 0xbd, 0x0f, 0xe0, 0x63, 0x59, 0xa6, 0xb4, 0x94, 0xe4, 0x11, 0x7c, 0x01, 0x31,
	0x77, 0x1b, 0x54, 0x22, 0x89, 0xe5, 0xcc, 0xfc, 0xdf, 0xfc, 0x33, 0xc3, 0xc0, 0x81, 0xca, 0x50,
	0x7b, 0xe5, 0x27, 0xe2, 0xb1, 0x9b, 0xa2, 0x97, 0x62, 0x5c, 0xa0, 0x53, 0xe8, 0x12, 0xeb, 0x8c,
	0x37, 0x51, 0x83, 0xac, 0x4e, 0x82, 0x24, 0x29, 0x25, 0xdd, 0x56, 0x67, 0x68, 0x28, 0x37, 0x24,
	0x52, 0x49, 0xb8, 0x00, 0x02, 0xdd, 0x15, 0x56, 0x8e, 0x94, 0x96, 0x5e, 0x19, 0x5d, 0x76, 0x88,
	0xcf, 0xe0, 0xf0, 0xe6, 0x4b, 0x71, 0x4b, 0xe8, 0xb4, 0xcc, 0xf1, 0xda, 0x99, 0xfc, 0x22, 0xcb,
	0x1c, 0x12, 0x0d, 0x70, 0x5c, 0x20, 0xf9, 0xa8, 0x09, 0x3b, 0xb2, 0xcc, 0x34, 0x59, 0x9b, 0x1d,
	0xed, 0x0e, 0x42, 0x18, 0x9f, 0x43, 0xfb, 0x6f, 0x98, 0xac, 0xd1, 0x84, 0x51, 0x0b, 0xea, 0x45,
	0x55, 0xae, 0xf0, 0x65, 0x1c, 0xf7, 0xa1, 0xf1, 0x83, 0x0f, 0x8e, 0x1c, 0xa0, 0x5c, 0xea, 0x5e,
	0xa1, 0xab, 0xa8, 0x6f, 0x99, 0xb8, 0x07, 0xfb, 0xbf, 0xb8, 0xf5, 0x66, 0x27, 0x1f, 0x0c, 0xb6,
	0x16, 0x54, 0xf4, 0xcc, 0x60, 0x6f, 0xc5, 0xc8, 0xd1, 0x69, 0xb2, 0xea, 0x9c, 0xc9, 0x9a, 0xfb,
	0xb4, 0xfa, 0xff, 0xc5, 0xaa, 0x61, 0x87, 0x50, 0x0f, 0xe5, 0xa8, 0xb3, 0x41, 0x8f, 0xe0, 0x77,
	0xbc, 0x91, 0xb6, 0x34, 0xb9, 0xbc, 0x7a, 0x9d, 0x71, 0x36, 0x9d, 0x71, 0xf6, 0x3e, 0xe3, 0xec,
	0x65, 0xce, 0x6b, 0xd3, 0x39, 0xaf, 0xbd, 0xcd, 0x79, 0xed, 0xae, 0x33, 0x52, 0xfe, 0xa1, 0x48,
	0x93, 0xa1, 0xc9, 0x85, 0x47, 0x9d, 0xa1, 0xcb, 0x95, 0xf6, 0x82, 0xac, 0x16, 0x4f, 0x62, 0xf9,
	0x76, 0x7e, 0x62, 0x91, 0xd2, 0xed, 0xc5, 0xb3, 0xf4, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xad,
	0x93, 0x23, 0x8c, 0x8f, 0x02, 0x00, 0x00,
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
	UsernameFromAddress(ctx context.Context, in *QueryUsernameFromAddressRequest, opts ...grpc.CallOption) (*QueryUsernameFromAddressResponse, error)
	Username(ctx context.Context, in *QueryUsernameRequest, opts ...grpc.CallOption) (*QueryUsernameResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) UsernameFromAddress(ctx context.Context, in *QueryUsernameFromAddressRequest, opts ...grpc.CallOption) (*QueryUsernameFromAddressResponse, error) {
	out := new(QueryUsernameFromAddressResponse)
	err := c.cc.Invoke(ctx, "/spn.identity.v1beta1.Query/UsernameFromAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Username(ctx context.Context, in *QueryUsernameRequest, opts ...grpc.CallOption) (*QueryUsernameResponse, error) {
	out := new(QueryUsernameResponse)
	err := c.cc.Invoke(ctx, "/spn.identity.v1beta1.Query/Username", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	UsernameFromAddress(context.Context, *QueryUsernameFromAddressRequest) (*QueryUsernameFromAddressResponse, error)
	Username(context.Context, *QueryUsernameRequest) (*QueryUsernameResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) UsernameFromAddress(ctx context.Context, req *QueryUsernameFromAddressRequest) (*QueryUsernameFromAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsernameFromAddress not implemented")
}
func (*UnimplementedQueryServer) Username(ctx context.Context, req *QueryUsernameRequest) (*QueryUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Username not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_UsernameFromAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUsernameFromAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UsernameFromAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spn.identity.v1beta1.Query/UsernameFromAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UsernameFromAddress(ctx, req.(*QueryUsernameFromAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Username_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Username(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spn.identity.v1beta1.Query/Username",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Username(ctx, req.(*QueryUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spn.identity.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UsernameFromAddress",
			Handler:    _Query_UsernameFromAddress_Handler,
		},
		{
			MethodName: "Username",
			Handler:    _Query_Username_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/v1beta/querier.proto",
}

func (m *QueryUsernameFromAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUsernameFromAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUsernameFromAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUsernameFromAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUsernameFromAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUsernameFromAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUsernameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUsernameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUsernameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUsernameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUsernameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUsernameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuerier(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuerier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryUsernameFromAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryUsernameFromAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryUsernameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryUsernameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func sovQuerier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuerier(x uint64) (n int) {
	return sovQuerier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryUsernameFromAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryUsernameFromAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUsernameFromAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryUsernameFromAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryUsernameFromAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUsernameFromAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryUsernameRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryUsernameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUsernameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryUsernameResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryUsernameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUsernameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuerier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func skipQuerier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
				return 0, ErrInvalidLengthQuerier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuerier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuerier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuerier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuerier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuerier = fmt.Errorf("proto: unexpected end of group")
)
