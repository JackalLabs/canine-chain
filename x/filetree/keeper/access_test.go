package keeper_test

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/google/uuid"

	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

func FuzzMakeViewerAddress(f *testing.F){
	bobPrivateK := secp256k1.GenPrivKey()
	bobPublicK := bobPrivateK.PubKey()
	hb := sha256.New()
	hb.Write(bobPublicK.Bytes())
	bobHash := fmt.Sprintf("%x", hb.Sum(nil))

	alicePrivateK := secp256k1.GenPrivKey()
	alicePublicK := alicePrivateK.PubKey()
	ha := sha256.New()
	ha.Write(alicePublicK.Bytes())
	aliceHash := fmt.Sprintf("%x", ha.Sum(nil))

	cases := []struct {
		trackingNum string
		user string
	}{
		{
			trackingNum: uuid.NewString(),
			user: aliceHash,
		},

		{
			trackingNum: uuid.NewString(),
			user: bobHash,
		},
	}

	for _, tc := range cases {
		f.Add(tc.trackingNum, tc.user)
	}

	f.Fuzz(func(t *testing.T, track, user string){
		out := keeper.MakeViewerAddress(track, user)

		eh := sha256.New()
		eh.Write([]byte(fmt.Sprintf("v%s%s", track, user)))
		expHash := fmt.Sprintf("%x", eh.Sum(nil))

		if out != expHash {
			t.Errorf("Expected: %s, Result: %s", expHash, out)
		}
	})
}

func (suite *KeeperTestSuite) TestMakeViewerAddress(){
	suite.SetupSuite()
	bobPrivateK := secp256k1.GenPrivKey()
	bobPublicK := bobPrivateK.PubKey()
	hb := sha256.New()
	hb.Write(bobPublicK.Bytes())
	bobHash := fmt.Sprintf("%x", hb.Sum(nil))

	alicePrivateK := secp256k1.GenPrivKey()
	alicePublicK := alicePrivateK.PubKey()
	ha := sha256.New()
	ha.Write(alicePublicK.Bytes())
	aliceHash := fmt.Sprintf("%x", ha.Sum(nil))


	cases := []struct {
		trackingNum string
		user string
	}{
		{
			trackingNum: uuid.NewString(),
			user: aliceHash,
		},

		{
			trackingNum: uuid.NewString(),
			user: bobHash,
		},
	}

	for _, tc := range cases {
		suite.reset()

		suite.Run(fmt.Sprintf("trackingNum: %s, user: %s", tc.trackingNum, tc.user), func() {
			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("v%s%s", tc.trackingNum, tc.user)))
			hash := h.Sum(nil)

			out := keeper.MakeViewerAddress(tc.trackingNum, tc.user)
			suite.Require().Equal(fmt.Sprintf("%x", hash), out)
		})
	}
}

func (suite *KeeperTestSuite) TestIsOwner(){
	suite.SetupSuite()

	cases := map[string]struct {
		addr string
		owner string
		user string
		expOut bool
	}{
		"true": {
			addr: "aaaaaaaa",
			owner: "",
			user: "alice",
			expOut: true,
		},

		"false": {
			addr: "aaaaaaaa",
			owner: "---------",
			user: "alice",
			expOut: false,
		},
	}

	for name, tc := range cases {
		suite.Run(name, func(){
			suite.reset()

			if tc.owner == "" {
				h := sha256.New()
				h.Write([]byte(tc.user))
				hash := h.Sum(nil)
				accountHash := fmt.Sprintf("%x", hash)

				// h1 is so named as to differentiate it from h above--else compiler complains
				h1 := sha256.New()
				h1.Write([]byte(fmt.Sprintf("o%s%s", tc.addr, accountHash)))
				hash1 := h1.Sum(nil)
				tc.owner = fmt.Sprintf("%x", hash1)
			}

			f := types.Files{
				Address: tc.addr,
				Owner: tc.owner,
			}

			out := keeper.IsOwner(f, tc.user)
			suite.Require().Equal(tc.expOut, out)
		})
	}
}

func FuzzHasEditAccess(f *testing.F) {
	cases := []struct {
		editAccess  string
		trackingNum string
		user        string
	}{
		{
			editAccess:  "aaaaaaaa",
			trackingNum: "111111111",
			user:        "",
		},

		{
			editAccess:  "",
			trackingNum: "someNum",
			user:        "someUser",
		},

		{
			editAccess:  `"diff_addr_str": "a"`,
			trackingNum: "someNum",
			user:        "someUser",
		},
	}

	for _, tc := range cases {
		f.Add(tc.editAccess, tc.trackingNum, tc.user)
	}
	
	f.Fuzz(func(t *testing.T, edit, track, user string){
		f := types.Files{
			EditAccess: edit,
			TrackingNumber: track,
		}

		out, err := keeper.HasEditAccess(f, user)

		if !json.Valid([]byte(edit)) {
			// The function should return error when invalid json is passed
			if err == nil {
				t.Errorf("Passed invalid json: %s but didn't get error.", edit)
			}

			h := sha256.New()
			h.Write([]byte(fmt.Sprintf("e%s%s", track, user)))
			hash := h.Sum(nil)

			jeacc := make(map[string]string)
			_ = json.Unmarshal([]byte(edit), &jeacc)
			_, expOut := jeacc[fmt.Sprintf("%x", hash)]

			if expOut != out {
				t.Errorf("Expected: %t, got: %t", expOut, out)
			}
		}
	})
}

func (suite *KeeperTestSuite) TestHasEditAccess() {
	suite.SetupSuite()

	cases := map[string]struct {
		editAccess  string
		trackingNum string
		user        string
		expErr      bool
		expResult   bool
	}{
		"invalid viewing access format": {
			editAccess:  "aaaaaaaa",
			trackingNum: "111111111",
			user:        "",
			expErr:      true,
			expResult:   false,
		},

		"has edit access": {
			editAccess:  "",
			trackingNum: "someNum",
			user:        "someUser",
			expErr:      false,
			expResult:   true,
		},

		"no edit access": {
			editAccess:  `"diff_addr_str": "a"`,
			trackingNum: "someNum",
			user:        "someUser",
			expErr:      false,
			expResult:   false,
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			if tc.editAccess == "" {
				// Construct new editor
				h := sha256.New()
				h.Write([]byte(fmt.Sprintf("e%s%s", tc.trackingNum, tc.user)))
				hash := h.Sum(nil)

				jeacc := make(map[string]string)
				jeacc[fmt.Sprintf("%x", hash)] = "a"
				eaccBytes, err := json.Marshal(jeacc)
				suite.Require().NoError(err)
				tc.editAccess = string(eaccBytes)
			}

			file := types.Files{
				EditAccess:     tc.editAccess,
				TrackingNumber: tc.trackingNum,
			}

			result, err := keeper.HasEditAccess(file, tc.user)

			if tc.expErr {
				suite.Require().Error(err)
			}
			suite.Require().Equal(tc.expResult, result)
		})
	}
}
