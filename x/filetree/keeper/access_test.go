package keeper_test

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jackalLabs/canine-chain/x/filetree/keeper"
	"github.com/jackalLabs/canine-chain/x/filetree/types"
)

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
