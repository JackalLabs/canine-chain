package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgPostContract_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPostContract
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgPostContract{
				Creator:  "invalid_address",
				Merkle:   "merkle",
				Signee:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Filesize: "100000",
				Fid:      "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid signee",
			msg: MsgPostContract{
				Creator:  "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Merkle:   "merkle",
				Signee:   "invalid_address",
				Filesize: "100000",
				Fid:      "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid filesize",
			msg: MsgPostContract{
				Creator:  "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Merkle:   "merkle",
				Signee:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Filesize: "x",
				Fid:      "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
			},
			err: sdkerrors.ErrInvalidType,
		}, {
			name: "invalid fid",
			msg: MsgPostContract{
				Creator:  "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Merkle:   "merkle",
				Signee:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Filesize: "100000",
				Fid:      "invalid_fid",
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address",
			msg: MsgPostContract{
				Creator:  "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Merkle:   "merkle",
				Signee:   "jkl1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Filesize: "100000",
				Fid:      "jklf1j3p63s42w7ywaczlju626st55mzu5z39w2rx9x",
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
