package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	apptypes "github.com/jackalLabs/canine-chain/types"
)

const TypeMsgCreateFeed = "create_feed"

var _ sdk.Msg = &MsgCreateFeed{}

func NewMsgCreateFeed(creator string, name string) *MsgCreateFeed {
	return &MsgCreateFeed{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgCreateFeed) Route() string {
	return RouterKey
}

func (msg *MsgCreateFeed) Type() string {
	return TypeMsgCreateFeed
}

func (msg *MsgCreateFeed) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateFeed) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateFeed) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != apptypes.Bech32Prefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}
