package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetProviderTotalSpace = "set_provider_total_space"

var _ sdk.Msg = &MsgSetProviderTotalSpace{}

func NewMsgSetProviderTotalSpace(creator string, space int64) *MsgSetProviderTotalSpace {
	return &MsgSetProviderTotalSpace{
		Creator: creator,
		Space:   space,
	}
}

func (msg *MsgSetProviderTotalSpace) Route() string {
	return RouterKey
}

func (msg *MsgSetProviderTotalSpace) Type() string {
	return TypeMsgSetProviderTotalSpace
}

func (msg *MsgSetProviderTotalSpace) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetProviderTotalSpace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetProviderTotalSpace) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}
