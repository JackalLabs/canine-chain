package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateFeed_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateFeed
		err  error
	}{
		{
			name: "valid feed",
			msg: MsgUpdateFeed{
				Creator: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Name:    "test feed 1",
			},
		},
		{
			name: "invalid creator",
			msg: MsgUpdateFeed{
				Creator: "cosmos1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Name:    "test feed 1",
			},
			err: sdkerrors.ErrInvalidAddress,
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
