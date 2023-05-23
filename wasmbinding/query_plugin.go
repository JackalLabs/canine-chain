package wasmbinding

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	bindings "github.com/jackalLabs/canine-chain/wasmbinding/bindings"
)

// CustomQuerier dispatches custom CosmWasm bindings queries.
func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.JackalQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, errorsmod.Wrap(err, "Jackal query")
		}

		switch {
		case contractQuery.Files != nil:
			address := contractQuery.Files.Address
			owner := contractQuery.Files.Owner

			// As it turns out, it's possible to safely use the keeper's GetFiles() method to return the Files struct
			// Using the query client is always still an option

			res, err := qp.GetFiles(ctx, address, owner)
			if err != nil {
				return nil, err
			}

			// !! WARNING !!
			// res is the Files struct converted to a one line string with the .String() method from files.pb.go
			// not sure if it's possible to json.Marshal this without an intermediate step
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, errorsmod.Wrap(err, "Jackal Files query response")
			}

			return bz, nil

		// can add more case statements
		// case contractQuery._________ != nil:

		default:
			return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown filetree query variant"}
		}
	}
}
