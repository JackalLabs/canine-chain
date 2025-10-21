package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	apptypes "github.com/jackalLabs/canine-chain/v5/types"
)

const TypeMsgUpdateFeed = "update_feed"

var _ sdk.Msg = &MsgUpdateFeed{}

func NewMsgUpdateFeed(creator string, name string, data string) *MsgUpdateFeed {
	return &MsgUpdateFeed{
		Creator: creator,
		Name:    name,
		Data:    data,
	}
}

func (msg *MsgUpdateFeed) Route() string {
	return RouterKey
}

func (msg *MsgUpdateFeed) Type() string {
	return TypeMsgUpdateFeed
}

func (msg *MsgUpdateFeed) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateFeed) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateFeed) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != apptypes.Bech32Prefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}
