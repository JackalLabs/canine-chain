package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgClaimStray_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgClaimStray
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgClaimStray{
				Creator:    "invalid_address",
				Cid:        "jklc1j3p63s42w7ywaczlju626st55mzu5z39qh6g4g",
				ForAddress: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid cid",
			msg: MsgClaimStray{
				Creator:    "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Cid:        "invalid_cid",
				ForAddress: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgClaimStray{
				Creator:    "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Cid:        "jklc1j3p63s42w7ywaczlju626st55mzu5z39qh6g4g",
				ForAddress: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
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
