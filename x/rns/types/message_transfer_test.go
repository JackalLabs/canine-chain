package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//	"github.com/jackalLabs/canine-chain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgTransfer
		err  error
	}{
		{
			name: "invalid sender",
			msg: MsgTransfer{
				Creator:  "invalid_address",
				Receiver: "cosmos1k3qu47ycrut4sr73vv6uqtuhyyfewymu34gju2",
				Name:     "validname.jkl",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid receiver",
			msg: MsgTransfer{
				Creator:  "cosmos1k3qu47ycrut4sr73vv6uqtuhyyfewymu34gju2",
				Receiver: "invalid_address",
				Name:     "validname.jkl",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid name",
			msg: MsgTransfer{
				Creator:  "cosmos1k3qu47ycrut4sr73vv6uqtuhyyfewymu34gju2",
				Receiver: "cosmos1k3qu47ycrut4sr73vv6uqtuhyyfewymu34gju2",
				Name:     "invalidname",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgTransfer{
				Creator:  "cosmos1k3qu47ycrut4sr73vv6uqtuhyyfewymu34gju2",
				Receiver: "cosmos1k3qu47ycrut4sr73vv6uqtuhyyfewymu34gju2",
				Name:     "validname.jkl",
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
