package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostZKProof = "postzkproof"

var _ sdk.Msg = &MsgPostZKProof{}

func NewMsgPostZKProof(creator string, hash string, zkPackage ProofPackage, hashlist string, cid string) *MsgPostZKProof {
	return &MsgPostZKProof{
		Creator:  creator,
		Cid:      cid,
		Package:  zkPackage,
		Hash:     hash,
		Hashlist: hashlist,
	}
}

func (msg *MsgPostZKProof) Route() string {
	return RouterKey
}

func (msg *MsgPostZKProof) Type() string {
	return TypeMsgPostZKProof
}

func (msg *MsgPostZKProof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostZKProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostZKProof) ValidateBasic() error {
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
