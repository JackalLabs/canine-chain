package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//	"github.com/jackalLabs/canine-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

// TODO: rewrite tests without ignite

func TestMsgList_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgList
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgList{
				Creator: "invalid_address",
				Name:    "validname.jkl",
				Price:   sdk.NewInt64Coin("ujkl", 1000),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid name",
			msg: MsgList{
				Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Name:    "invalidname",
				Price:   sdk.NewInt64Coin("ujkl", 1000),
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgList{
				Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Name:    "validname.jkl",
				Price:   sdk.NewInt64Coin("ujkl", 1000),
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
