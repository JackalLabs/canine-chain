package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgPostProof_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPostProof
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgPostProof{
				Creator:  "invalid_address",
				Item:     "hex",
				HashList: []byte{},
				Merkle:   []byte{},
				Owner:    "owner",
				Start:    0,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgPostProof{
				Creator:  "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Item:     "hex",
				HashList: []byte{},
				Merkle:   []byte{},
				Owner:    "owner",
				Start:    0,
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
