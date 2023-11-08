package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgInitProvider_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgInitProvider
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgInitProvider{
				Creator:    "invalid_address",
				Ip:         "http://localhost:3333",
				TotalSpace: "1000000000",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid ip",
			msg: MsgInitProvider{
				Creator:    "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Ip:         "fake/localhost:3333",
				TotalSpace: "1000000000",
			},
			err: sdkerrors.ErrInvalidType,
		}, {
			name: "invalid totalspace",
			msg: MsgInitProvider{
				Creator:    "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Ip:         "https://node.jackalprotocol.com",
				TotalSpace: "abd",
			},
			err: sdkerrors.ErrInvalidType,
		}, {
			name: "valid ip",
			msg: MsgInitProvider{
				Creator:    "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Ip:         "https://node.jackalprotocol.com",
				TotalSpace: "1000000000",
			},
		}, {
			name: "valid ip localhost",
			msg: MsgInitProvider{
				Creator:    "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Ip:         "localhost:3333",
				TotalSpace: "1000000000",
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
