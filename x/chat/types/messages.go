package types

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
)

// Chat message types
const (
	TypeMsgCreateChannel = "create_channel"
	TypeMsgSendMessage   = "send_message"
	TypeMsgVotePoll      = "vote_poll"
)

// Verify interface at compile time
var (
	_ sdk.Msg = &MsgCreateChannel{}
	_ sdk.Msg = &MsgSendMessage{}
	_ sdk.Msg = &MsgVotePoll{}
)

// MsgCreateChannel

// NewMsgCreateChannel creates a new message to create a channel
func NewMsgCreateChannel(
	creator User,
	name string,
	subject string,
	metadata proto.Message,
) (*MsgCreateChannel, error) {
	// Convert the proto message in any
	metadataAny, err := types.NewAnyWithValue(metadata)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidChannel, err.Error())
	}

	return &MsgCreateChannel{
		Creator:  &creator,
		Name:     name,
		Subject:  subject,
		Metadata: metadataAny,
	}, nil
}

// Route implements the sdk.Msg interface.
func (msg MsgCreateChannel) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgCreateChannel) Type() string { return TypeMsgCreateChannel }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgCreateChannel) GetSigners() []sdk.AccAddress {
	// Get the address from the creator
	addressableUser, err := msg.Creator.DecodeAddressableUser()

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addressableUser.Address()}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgCreateChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCreateChannel) ValidateBasic() error {
	// TODO: Message validate basics
	return nil
}

// MsgSendMessage

// NewMsgSendMessage creates a new message to send a message to a chanel
func NewMsgSendMessage(
	channelID int32,
	author User,
	content string,
	tags []string,
	pollOptions []string,
	metadata proto.Message,
) (*MsgSendMessage, error) {
	// Convert the proto message in any
	metadataAny, err := types.NewAnyWithValue(metadata)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidChannel, err.Error())
	}

	return &MsgSendMessage{
		ChannelID:   channelID,
		Author:      &author,
		Content:     content,
		Tags:        tags,
		PollOptions: pollOptions,
		Metadata:    metadataAny,
	}, nil
}

// Route implements the sdk.Msg interface.
func (msg MsgSendMessage) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgSendMessage) Type() string { return TypeMsgSendMessage }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgSendMessage) GetSigners() []sdk.AccAddress {
	// Get the address from the creator
	addressableUser, err := msg.Author.DecodeAddressableUser()

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addressableUser.Address()}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgSendMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgSendMessage) ValidateBasic() error {
	// TODO: Message validate basics
	return nil
}

// MsgVotePoll

// NewMsgVotePoll creates a new message to vote to a poll
func NewMsgVotePoll(
	messageID string,
	author User,
	value int32,
	metadata proto.Message,
) (*MsgVotePoll, error) {
	// Convert the proto message in any
	metadataAny, err := types.NewAnyWithValue(metadata)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidChannel, err.Error())
	}

	return &MsgVotePoll{
		MessageID: messageID,
		Author:    &author,
		Value:     value,
		Metadata:  metadataAny,
	}, nil
}

// Route implements the sdk.Msg interface.
func (msg MsgVotePoll) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgVotePoll) Type() string { return TypeMsgVotePoll }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgVotePoll) GetSigners() []sdk.AccAddress {
	// Get the address from the creator
	addressableUser, err := msg.Author.DecodeAddressableUser()

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addressableUser.Address()}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgVotePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgVotePoll) ValidateBasic() error {
	// TODO: Message validate basics
	return nil
}
