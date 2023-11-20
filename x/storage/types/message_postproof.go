package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostProof = "post_proof"

var _ sdk.Msg = &MsgPostProof{}

func NewMsgPostProof(creator string, merkle []byte, owner string, start int64, item []byte, hashList []byte, toProve int64) *MsgPostProof {
	return &MsgPostProof{
		Creator:  creator,
		Item:     item,
		HashList: hashList,
		Merkle:   merkle,
		Owner:    owner,
		Start:    start,
		ToProve:  toProve,
	}
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
