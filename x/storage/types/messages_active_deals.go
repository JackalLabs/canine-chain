package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateActiveDeals = "create_active_deals"
	TypeMsgUpdateActiveDeals = "update_active_deals"
	TypeMsgDeleteActiveDeals = "delete_active_deals"
)

var _ sdk.Msg = &MsgCreateActiveDeals{}

func NewMsgCreateActiveDeals(
	creator string,
	cid string,
	signee string,
	provider string,
	startblock string,
	endblock string,
	filesize string,
	proofverified string,
	proofsmissed string,
	blocktoprove string,
) *MsgCreateActiveDeals {
	return &MsgCreateActiveDeals{
		Creator:       creator,
		Cid:           cid,
		Signee:        signee,
		Provider:      provider,
		Startblock:    startblock,
		Endblock:      endblock,
		Filesize:      filesize,
		Proofverified: proofverified,
		Proofsmissed:  proofsmissed,
		Blocktoprove:  blocktoprove,
	}
}

func (msg *MsgCreateActiveDeals) Route() string {
	return RouterKey
}

func (msg *MsgCreateActiveDeals) Type() string {
	return TypeMsgCreateActiveDeals
}

func (msg *MsgCreateActiveDeals) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateActiveDeals) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateActiveDeals) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateActiveDeals{}

func NewMsgUpdateActiveDeals(
	creator string,
	cid string,
	signee string,
	provider string,
	startblock string,
	endblock string,
	filesize string,
	proofverified string,
	proofsmissed string,
	blocktoprove string,
) *MsgUpdateActiveDeals {
	return &MsgUpdateActiveDeals{
		Creator:       creator,
		Cid:           cid,
		Signee:        signee,
		Provider:      provider,
		Startblock:    startblock,
		Endblock:      endblock,
		Filesize:      filesize,
		Proofverified: proofverified,
		Proofsmissed:  proofsmissed,
		Blocktoprove:  blocktoprove,
	}
}

func (msg *MsgUpdateActiveDeals) Route() string {
	return RouterKey
}

func (msg *MsgUpdateActiveDeals) Type() string {
	return TypeMsgUpdateActiveDeals
}

func (msg *MsgUpdateActiveDeals) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateActiveDeals) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateActiveDeals) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteActiveDeals{}

func NewMsgDeleteActiveDeals(
	creator string,
	cid string,
) *MsgDeleteActiveDeals {
	return &MsgDeleteActiveDeals{
		Creator: creator,
		Cid:     cid,
	}
}

func (msg *MsgDeleteActiveDeals) Route() string {
	return RouterKey
}

func (msg *MsgDeleteActiveDeals) Type() string {
	return TypeMsgDeleteActiveDeals
}

func (msg *MsgDeleteActiveDeals) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteActiveDeals) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteActiveDeals) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
