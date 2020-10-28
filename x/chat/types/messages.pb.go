// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat/v1beta/messages.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type MsgCreateChannel struct {
	Creator github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Name    string                                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Subject string                                        `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Payload *types.Any                                    `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgCreateChannel) Reset()         { *m = MsgCreateChannel{} }
func (m *MsgCreateChannel) String() string { return proto.CompactTextString(m) }
func (*MsgCreateChannel) ProtoMessage()    {}
func (*MsgCreateChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_3395017881b3c0cc, []int{0}
}
func (m *MsgCreateChannel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateChannel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateChannel.Merge(m, src)
}
func (m *MsgCreateChannel) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateChannel.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateChannel proto.InternalMessageInfo

func (m *MsgCreateChannel) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgCreateChannel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCreateChannel) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *MsgCreateChannel) GetPayload() *types.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type MsgSendMessage struct {
	ChannelID   int32                                         `protobuf:"varint,1,opt,name=channelID,proto3" json:"channelID,omitempty"`
	Creator     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Content     string                                        `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Tags        []string                                      `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	PollOptions []string                                      `protobuf:"bytes,5,rep,name=pollOptions,proto3" json:"pollOptions,omitempty"`
	Payload     *types.Any                                    `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgSendMessage) Reset()         { *m = MsgSendMessage{} }
func (m *MsgSendMessage) String() string { return proto.CompactTextString(m) }
func (*MsgSendMessage) ProtoMessage()    {}
func (*MsgSendMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3395017881b3c0cc, []int{1}
}
func (m *MsgSendMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendMessage.Merge(m, src)
}
func (m *MsgSendMessage) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendMessage proto.InternalMessageInfo

func (m *MsgSendMessage) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *MsgSendMessage) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgSendMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MsgSendMessage) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MsgSendMessage) GetPollOptions() []string {
	if m != nil {
		return m.PollOptions
	}
	return nil
}

func (m *MsgSendMessage) GetPayload() *types.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type MsgVotePoll struct {
	MessageID string                                        `protobuf:"bytes,1,opt,name=messageID,proto3" json:"messageID,omitempty"`
	Creator   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Value     int32                                         `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	Payload   *types.Any                                    `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgVotePoll) Reset()         { *m = MsgVotePoll{} }
func (m *MsgVotePoll) String() string { return proto.CompactTextString(m) }
func (*MsgVotePoll) ProtoMessage()    {}
func (*MsgVotePoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_3395017881b3c0cc, []int{2}
}
func (m *MsgVotePoll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVotePoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVotePoll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVotePoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVotePoll.Merge(m, src)
}
func (m *MsgVotePoll) XXX_Size() int {
	return m.Size()
}
func (m *MsgVotePoll) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVotePoll.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVotePoll proto.InternalMessageInfo

func (m *MsgVotePoll) GetMessageID() string {
	if m != nil {
		return m.MessageID
	}
	return ""
}

func (m *MsgVotePoll) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgVotePoll) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MsgVotePoll) GetPayload() *types.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCreateChannel)(nil), "spn.chat.v1beta1.MsgCreateChannel")
	proto.RegisterType((*MsgSendMessage)(nil), "spn.chat.v1beta1.MsgSendMessage")
	proto.RegisterType((*MsgVotePoll)(nil), "spn.chat.v1beta1.MsgVotePoll")
}

func init() { proto.RegisterFile("chat/v1beta/messages.proto", fileDescriptor_3395017881b3c0cc) }

var fileDescriptor_3395017881b3c0cc = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xeb, 0xdd, 0x76, 0xab, 0xba, 0x08, 0xad, 0xa2, 0x0a, 0x85, 0x0a, 0x85, 0x68, 0x0f,
	0xa8, 0x97, 0xb5, 0x55, 0x78, 0x01, 0xba, 0xcb, 0x05, 0xa1, 0x0a, 0x14, 0x24, 0x0e, 0xdc, 0x1c,
	0x67, 0x70, 0x0b, 0x8e, 0x1d, 0x65, 0xdc, 0x15, 0x7d, 0x0b, 0x5e, 0x08, 0xce, 0x1c, 0xf7, 0xc8,
	0x09, 0xa1, 0xf6, 0x2d, 0xe0, 0x82, 0x62, 0x37, 0x6c, 0x8e, 0x20, 0xf6, 0x94, 0x7f, 0xfe, 0x7f,
	0x32, 0x9a, 0x2f, 0x19, 0x3a, 0x95, 0x2b, 0xe1, 0xf8, 0xd5, 0x3c, 0x07, 0x27, 0x78, 0x09, 0x88,
	0x42, 0x01, 0xb2, 0xaa, 0xb6, 0xce, 0x46, 0xa7, 0x58, 0x19, 0xd6, 0xe4, 0x2c, 0xe4, 0xf3, 0xe9,
	0x7d, 0x65, 0xad, 0xd2, 0xc0, 0x7d, 0x9e, 0x6f, 0xde, 0x71, 0x61, 0xb6, 0xa1, 0x79, 0x7a, 0xaf,
	0x3b, 0xc8, 0xbf, 0x14, 0xfc, 0x89, 0xb2, 0xca, 0x7a, 0xc9, 0x1b, 0x15, 0xdc, 0xb3, 0xcf, 0x84,
	0x9e, 0x2e, 0x51, 0x5d, 0xd6, 0x20, 0x1c, 0x5c, 0xae, 0x84, 0x31, 0xa0, 0xa3, 0x17, 0x74, 0x28,
	0x1b, 0xc3, 0xd6, 0x31, 0x49, 0xc9, 0xec, 0xce, 0xc5, 0xfc, 0xe7, 0xf7, 0x87, 0xe7, 0x6a, 0xed,
	0x56, 0x9b, 0x9c, 0x49, 0x5b, 0x72, 0x69, 0xb1, 0xb4, 0x78, 0x78, 0x9c, 0x63, 0xf1, 0x81, 0xbb,
	0x6d, 0x05, 0xc8, 0x16, 0x52, 0x2e, 0x8a, 0xa2, 0x06, 0xc4, 0xac, 0x9d, 0x10, 0x45, 0xb4, 0x6f,
	0x44, 0x09, 0xf1, 0x51, 0x4a, 0x66, 0xa3, 0xcc, 0xeb, 0x28, 0xa6, 0x43, 0xdc, 0xe4, 0xef, 0x41,
	0xba, 0xf8, 0xd8, 0xdb, 0x6d, 0x19, 0x31, 0x3a, 0xac, 0xc4, 0x56, 0x5b, 0x51, 0xc4, 0xfd, 0x94,
	0xcc, 0xc6, 0x8f, 0x27, 0x2c, 0xa0, 0xb2, 0x16, 0x95, 0x2d, 0xcc, 0x36, 0x6b, 0x9b, 0xce, 0x7e,
	0x11, 0x7a, 0x77, 0x89, 0xea, 0x35, 0x98, 0x62, 0x19, 0x3e, 0x5a, 0xf4, 0x80, 0x8e, 0x64, 0x00,
	0x79, 0xfe, 0xcc, 0xef, 0x3f, 0xc8, 0x6e, 0x8c, 0x2e, 0xdb, 0xd1, 0x7f, 0xb3, 0xc5, 0x74, 0x28,
	0xad, 0x71, 0x60, 0xfe, 0x70, 0x1c, 0xca, 0x86, 0xda, 0x09, 0x85, 0x71, 0x3f, 0x3d, 0x6e, 0xa8,
	0x1b, 0x1d, 0xa5, 0x74, 0x5c, 0x59, 0xad, 0x5f, 0x56, 0x6e, 0x6d, 0x0d, 0xc6, 0x03, 0x1f, 0x75,
	0xad, 0x2e, 0xfd, 0xc9, 0xdf, 0xd0, 0x7f, 0x21, 0x74, 0xbc, 0x44, 0xf5, 0xc6, 0x3a, 0x78, 0x65,
	0xb5, 0x6e, 0xd0, 0x0f, 0xa7, 0x73, 0x40, 0x1f, 0x65, 0x37, 0xc6, 0xed, 0xa2, 0x4f, 0xe8, 0xe0,
	0x4a, 0xe8, 0x0d, 0x78, 0xf0, 0x41, 0x16, 0x8a, 0x7f, 0xfd, 0x7d, 0x17, 0x4f, 0xbf, 0xee, 0x12,
	0x72, 0xbd, 0x4b, 0xc8, 0x8f, 0x5d, 0x42, 0x3e, 0xed, 0x93, 0xde, 0xf5, 0x3e, 0xe9, 0x7d, 0xdb,
	0x27, 0xbd, 0xb7, 0x8f, 0x3a, 0x7b, 0x39, 0x30, 0x05, 0xd4, 0xe5, 0xda, 0x38, 0x8e, 0x95, 0xe1,
	0x1f, 0xfd, 0x59, 0x87, 0xdd, 0xf2, 0x13, 0x3f, 0xf8, 0xc9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x55, 0x70, 0x14, 0x46, 0x40, 0x03, 0x00, 0x00,
}

func (m *MsgCreateChannel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateChannel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateChannel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		{
			size, err := m.Payload.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessages(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		{
			size, err := m.Payload.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessages(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.PollOptions) > 0 {
		for iNdEx := len(m.PollOptions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PollOptions[iNdEx])
			copy(dAtA[i:], m.PollOptions[iNdEx])
			i = encodeVarintMessages(dAtA, i, uint64(len(m.PollOptions[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintMessages(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChannelID != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.ChannelID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVotePoll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVotePoll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVotePoll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		{
			size, err := m.Payload.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessages(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Value != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MessageID) > 0 {
		i -= len(m.MessageID)
		copy(dAtA[i:], m.MessageID)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.MessageID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateChannel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgSendMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChannelID != 0 {
		n += 1 + sovMessages(uint64(m.ChannelID))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	if len(m.PollOptions) > 0 {
		for _, s := range m.PollOptions {
			l = len(s)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgVotePoll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MessageID)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovMessages(uint64(m.Value))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateChannel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgCreateChannel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateChannel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &types.Any{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgSendMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSendMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			m.ChannelID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChannelID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollOptions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PollOptions = append(m.PollOptions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &types.Any{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgVotePoll) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgVotePoll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVotePoll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &types.Any{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessages
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
