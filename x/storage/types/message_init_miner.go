package types

import (
	fmt "fmt"
	"net/url"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitProvider = "init_provider"

var _ sdk.Msg = &MsgInitProvider{}

func NewMsgInitProvider(creator string, ip string, totalspace string, keybase string) *MsgInitProvider {
	return &MsgInitProvider{
		Creator:    creator,
		Ip:         ip,
		Totalspace: totalspace,
		Keybase:    keybase,
	}
}

func (msg *MsgInitProvider) Route() string {
	return RouterKey
}

func (msg *MsgInitProvider) Type() string {
	return TypeMsgInitProvider
}

func (msg *MsgInitProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitProvider) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	_, err = url.ParseRequestURI(msg.Ip)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid provider ip (%s)", err)
	}

	if _, err := strconv.Atoi(msg.Totalspace); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse totalspace (%s)", err)
	}
	return nil
}
