package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSetProviderKeybase_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetProviderKeybase
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetProviderKeybase{
				Creator: "invalid_address",
				Keybase: "test-key",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid key & address",
			msg: MsgSetProviderKeybase{
				Creator: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Keybase: "test-key",
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
