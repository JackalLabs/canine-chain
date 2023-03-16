package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestMsgPostFile_ValidateBasic(t *testing.T) {
	alicePrivateK := secp256k1.GenPrivKey()
	alicePublicK := alicePrivateK.PubKey()
	aliceAddr := sdk.AccAddress(alicePublicK.Address())

	tests := map[string]struct {
		Creator, Account, HashParent, HashChild, Contents, Viewers, Editors, TrackingNumber string
		expErr                                                                              bool
	}{
		"invalid address": {
			Creator:        "",
			Account:        uuid.NewString(),
			HashParent:     uuid.NewString(),
			HashChild:      uuid.NewString(),
			Contents:       "",
			Viewers:        uuid.NewString(),
			Editors:        uuid.NewString(),
			TrackingNumber: uuid.NewString(),
			expErr:         true,
		},

		"valid address": {
			Creator:        aliceAddr.String(),
			Account:        uuid.NewString(),
			HashParent:     uuid.NewString(),
			HashChild:      uuid.NewString(),
			Contents:       "",
			Viewers:        aliceAddr.String(),
			Editors:        uuid.NewString(),
			TrackingNumber: uuid.NewString(),
			expErr:         false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			msg := NewMsgPostFile(
				tt.Creator, tt.Account, tt.HashParent, tt.HashChild, tt.Contents, tt.Viewers, tt.Editors, tt.TrackingNumber,
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
