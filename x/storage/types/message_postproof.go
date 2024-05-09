package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgPostproof = "postproof"
	TypeMsgPostProof = "post_proof"
)

var (
	_ sdk.Msg = &MsgPostproof{}
	_ sdk.Msg = &MsgPostProof{}
)

func NewMsgPostproof(creator string, item string, hashlist string, cid string) *MsgPostproof {
	return &MsgPostproof{
		Creator:  creator,
		Cid:      cid,
		Item:     item,
		Hashlist: hashlist,
	}
}

func NewMsgPostProof(creator string, merkle []byte, owner string, start int64, item []byte, proof []byte, index int64) *MsgPostProof {
	return &MsgPostProof{
		Creator:  creator,
		Item:     item,
		HashList: proof,
		Merkle:   merkle,
		Owner:    owner,
		Start:    start,
		ToProve:  index,
	}
}

func (msg *MsgPostproof) Route() string {
	return RouterKey
}

func (msg *MsgPostproof) Type() string {
	return TypeMsgPostproof
}

func (msg *MsgPostproof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostproof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostproof) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	prefix, _, err = bech32.DecodeAndConvert(msg.Cid)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid cid (%s)", err)
	}
	if prefix != "jklc" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid cid prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jklc`", prefix))
	}
	return nil
}

func (msg *MsgPostProof) Route() string {
	return RouterKey
}

func (msg *MsgPostProof) Type() string {
	return TypeMsgPostProof
}

func (msg *MsgPostProof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostProof) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}
