package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRequestChunk   = "request_oracle_chunk"
	TypeMsgFulfillRequest = "fulfill_oracle_request"
)

var (
	_ sdk.Msg = &MsgRequestChunk{}
	_ sdk.Msg = &MsgFulfillRequest{}
)

func NewMsgRequestChunk(creator string, merkle []byte, chunk int64, bid sdk.Coin) *MsgRequestChunk {
	return &MsgRequestChunk{
		Creator: creator,
		Merkle:  merkle,
		Chunk:   chunk,
		Bid:     bid,
	}
}

func (msg *MsgRequestChunk) Route() string {
	return RouterKey
}

func (msg *MsgRequestChunk) Type() string {
	return TypeMsgRequestChunk
}

func (msg *MsgRequestChunk) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestChunk) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestChunk) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}

func NewMsgFulfillRequest(creator string, merkle []byte, requester string, chunk int64, data []byte) *MsgFulfillRequest {
	return &MsgFulfillRequest{
		Creator:   creator,
		Merkle:    merkle,
		Requester: requester,
		Chunk:     chunk,
		Data:      data,
	}
}

func (msg *MsgFulfillRequest) Route() string {
	return RouterKey
}

func (msg *MsgFulfillRequest) Type() string {
	return TypeMsgFulfillRequest
}

func (msg *MsgFulfillRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFulfillRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFulfillRequest) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}
