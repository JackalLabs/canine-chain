package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgBuyStorage_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgBuyStorage
		err  error
	}{
		{
			name: "invalid for address",
			msg: MsgBuyStorage{
				Creator:      "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				ForAddress:   "invalid_address",
				DurationDays: 30,
				Bytes:        4096,
				PaymentDenom: "ujkl",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid from address",
			msg: MsgBuyStorage{
				Creator:      "invalid_address",
				ForAddress:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				DurationDays: 30,
				Bytes:        4096,
				PaymentDenom: "ujkl",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid duration",
			msg: MsgBuyStorage{
				Creator:      "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				ForAddress:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				DurationDays: 0,
				Bytes:        4096,
				PaymentDenom: "ujkl",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid duration(negative)",
			msg: MsgBuyStorage{
				Creator:      "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				ForAddress:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				DurationDays: -10,
				Bytes:        4096,
				PaymentDenom: "ujkl",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgBuyStorage{
				Creator:      "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				ForAddress:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				DurationDays: 30,
				Bytes:        4096,
				PaymentDenom: "ujkl",
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
