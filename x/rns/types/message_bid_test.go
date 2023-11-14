package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//	"github.com/jackalLabs/canine-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgBid_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgBid
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgBid{
				Creator: "invalid_address",
				Name:    "validname.jkl",
				Bid:     sdk.NewInt64Coin("ujkl", 10000),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid name",
			msg: MsgBid{
				Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Name:    "invalidname",
				Bid:     sdk.NewInt64Coin("ujkl", 1000),
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgBid{
				Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Name:    "validname.jkl",
				Bid:     sdk.NewInt64Coin("ujkl", 1000),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
