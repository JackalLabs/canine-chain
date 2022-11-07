package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateProviders_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateProviders
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateProviders{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg:  MsgCreateProviders{
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

func TestMsgUpdateProviders_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateProviders
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateProviders{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg:  MsgUpdateProviders{
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

func TestMsgDeleteProviders_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteProviders
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteProviders{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg:  MsgDeleteProviders{
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
