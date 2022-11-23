package keeper_test

import (
//	"testing"

	"github.com/jackalLabs/canine-chain/x/filetree/types"
//	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
//	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) HasEditAccess(){

	cases := map[string]struct{
		editAccess string
		trackingNum string
		user string 
		expResult bool 
	}{
		"invalid viewing access format": {
			editAccess: "aaaaaaaa",
			trackingNum: "111111111",
			user: "",
			expResult: false,
		},
	}

	for name, tc := range cases {
		suite.Run(name, func(){

			file := types.Files{
				EditAccess: tc.editAccess,
				TrackingNumber: tc.trackingNum,
			}
			result := keeper.HasEditAccess(file, tc.user)

			suite.Require().Equal(tc.expResult, result)
		})
	}

}
