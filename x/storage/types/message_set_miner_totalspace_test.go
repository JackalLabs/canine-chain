package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSetProviderTotalSpace_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetProviderTotalSpace
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetProviderTotalSpace{
				Creator: "invalid_address",
				Space:   "1000000000",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid ip",
			msg: MsgSetProviderTotalSpace{
				Creator: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Space:   "abd",
			},
			err: sdkerrors.ErrInvalidType,
		}, {
			name: "valid ip",
			msg: MsgSetProviderTotalSpace{
				Creator: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Space:   "1000000000",
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
