package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//	"github.com/jackalLabs/canine-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgAddRecord_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddRecord
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAddRecord{
				Creator: "invalid_address",
				Name:    "validname.jkl",
				Value:   "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Data:    "{}",
				Record:  "app",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid name",
			msg: MsgAddRecord{
				Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Name:    "invalidname",
				Value:   "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Data:    "{}",
				Record:  "app",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgAddRecord{
				Creator: "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Name:    "validname.jkl",
				Value:   "cosmos1ytwr7x4av05ek0tf8z9s4zmvr6w569zsm27dpg",
				Data:    "{}",
				Record:  "app",
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
