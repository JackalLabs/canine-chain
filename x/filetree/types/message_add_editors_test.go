package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAddEditors_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddEditors
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgAddEditors{
				Creator:   "invalid_address",
				Fileowner: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid fileowner address",
			msg: MsgAddEditors{
				Creator:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Fileowner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid creator, valid fileowner",
			msg: MsgAddEditors{
				Creator:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Fileowner: "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
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
