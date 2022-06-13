package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateMiners_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateMiners
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateMiners{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateMiners{
				Creator: sample.AccAddress(),
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

func TestMsgUpdateMiners_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateMiners
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateMiners{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateMiners{
				Creator: sample.AccAddress(),
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

func TestMsgDeleteMiners_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteMiners
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteMiners{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteMiners{
				Creator: sample.AccAddress(),
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
