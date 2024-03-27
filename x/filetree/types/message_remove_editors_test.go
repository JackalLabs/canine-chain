package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestMsgRemoveEditors_ValidateBasic(t *testing.T) {
	alicePrivateK := secp256k1.GenPrivKey()
	alicePublicK := alicePrivateK.PubKey()
	aliceAddr := sdk.AccAddress(alicePublicK.Address())

	tests := map[string]struct {
		Creator, EditorIDs, Address, Fileowner string
		expErr                                 bool
	}{
		"invalid address": {
			Creator:   "",
			Address:   uuid.NewString(),
			EditorIDs: uuid.NewString(),
			Fileowner: uuid.NewString(),
			expErr:    true,
		},

		"valid address": {
			Creator:   aliceAddr.String(),
			Address:   uuid.NewString(),
			EditorIDs: uuid.NewString(),
			Fileowner: uuid.NewString(),
			expErr:    false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			msg := NewMsgRemoveEditors(
				tt.Creator, tt.Address, tt.EditorIDs, tt.Fileowner,
			)

			err := msg.ValidateBasic()
			t.Logf("Address: %s", msg.Creator)
			if tt.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
