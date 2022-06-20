package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackal-dao/canine/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateSaveRequests_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateSaveRequests
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSaveRequests{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSaveRequests{
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

func TestMsgUpdateSaveRequests_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSaveRequests
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSaveRequests{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSaveRequests{
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

func TestMsgDeleteSaveRequests_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteSaveRequests
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteSaveRequests{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteSaveRequests{
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
