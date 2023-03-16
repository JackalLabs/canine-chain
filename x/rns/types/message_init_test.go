package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//	"github.com/jackalLabs/canine-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

// TODO: rewrite tests without ignite

func TestMsgInit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgInit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgInit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgInit{
				Creator: "cosmos1k3qu47ycrut4sr73vv6uqtuhyyfewymu34gju2",
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
