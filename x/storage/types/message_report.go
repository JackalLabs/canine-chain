package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRequestReportForm = "request_report_form"
	TypeMsgReport            = "report"
)

var (
	_ sdk.Msg = &MsgRequestReportForm{}
	_ sdk.Msg = &MsgReport{}
)

func NewMsgRequestReportForm(creator string, prover string, merkle []byte, owner string, start int64) *MsgRequestReportForm {
	return &MsgRequestReportForm{
		Creator: creator,
		Prover:  prover,
		Merkle:  merkle,
		Owner:   owner,
		Start:   start,
	}
}

func (msg *MsgRequestReportForm) Route() string {
	return RouterKey
}

func (msg *MsgRequestReportForm) Type() string {
	return TypeMsgRequestReportForm
}

func (msg *MsgRequestReportForm) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestReportForm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestReportForm) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}

func NewMsgReport(creator string, prover string, merkle []byte, owner string, start int64) *MsgReport {
	return &MsgReport{
		Creator: creator,
		Prover:  prover,
		Merkle:  merkle,
		Owner:   owner,
		Start:   start,
	}
}

func (msg *MsgReport) Route() string {
	return RouterKey
}

func (msg *MsgReport) Type() string {
	return TypeMsgReport
}

func (msg *MsgReport) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReport) ValidateBasic() error {
	prefix, _, err := bech32.DecodeAndConvert(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if prefix != AddressPrefix {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator prefix (%s)", fmt.Errorf("%s is not a valid prefix here. Expected `jkl`", prefix))
	}

	return nil
}
