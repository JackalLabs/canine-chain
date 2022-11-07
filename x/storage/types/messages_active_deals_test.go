package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateActiveDeals_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateActiveDeals
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateActiveDeals{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg:  MsgCreateActiveDeals{
				// Creator: sample.AccAddress(),
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

func TestMsgUpdateActiveDeals_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateActiveDeals
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateActiveDeals{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg:  MsgUpdateActiveDeals{
				// Creator: sample.AccAddress(),
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

func TestMsgDeleteActiveDeals_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteActiveDeals
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteActiveDeals{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg:  MsgDeleteActiveDeals{
				// Creator: sample.AccAddress(),
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
