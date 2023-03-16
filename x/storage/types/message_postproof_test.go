package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgPostproof_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPostproof
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgPostproof{
				Creator:  "invalid_address",
				Item:     "hex",
				Hashlist: "hex",
				Cid:      "jklc1j3p63s42w7ywaczlju626st55mzu5z39qh6g4g",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid cid",
			msg: MsgPostproof{
				Creator:  "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Item:     "hex",
				Hashlist: "hex",
				Cid:      "invalid_cid",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgPostproof{
				Creator:  "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Item:     "hex",
				Hashlist: "hex",
				Cid:      "jklc1j3p63s42w7ywaczlju626st55mzu5z39qh6g4g",
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
