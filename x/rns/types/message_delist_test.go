package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//	"github.com/jackalLabs/canine-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgDelist_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDelist
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDelist{
				Creator: "invalid_address",
				Name:    "validname.jkl",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid name",
			msg: MsgDelist{
				Creator: "cosmos1k3qu47ycrut4sr73vv6uqtuhyyfewymu34gju2",
				Name:    "invalidname",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgDelist{
				Creator: "cosmos1k3qu47ycrut4sr73vv6uqtuhyyfewymu34gju2",
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
