package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//	"github.com/jackalLabs/canine-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCancelBid_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelBid
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCancelBid{
				Creator: "invalid_address",
				Name:    "validname.jkl",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid name",
			msg: MsgCancelBid{
				Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Name:    "invalidname",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgCancelBid{
				Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Name:    "validname.jkl",
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
