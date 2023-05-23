package wasmbinding

import (
	"context"
	"encoding/json"

	grpc1 "github.com/gogo/protobuf/grpc"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	bindings "github.com/jackalLabs/canine-chain/wasmbinding/bindings"
	filetreetypes "github.com/jackalLabs/canine-chain/x/filetree/types"
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
			owner := contractQuery.Files.Owner
			address := contractQuery.Files.Address

			// Osmosis uses GetFullDenom which calls tokenfactorytypes.GetTokenDenom(contract, subDenom)
			// which simply just builds a string. This is not useful to us.
			// We can fork wasmd to expose the keeper, or expose the keeper some other way, but this
			// is a huge security vulnerability.
			// Perhaps we can just use the query client here? If this package is not a full-on module, does it slow the chain down at all?

			// Don't need the below code any more
			// Files, err := RetrieveFiles(ctx, filetree, owner, address)

			var cc grpc1.ClientConn

			queryClient := filetreetypes.NewQueryClient(cc)

			params := &filetreetypes.QueryFileRequest{
				Address:      address,
				OwnerAddress: owner,
			}

			queryRes, err := queryClient.Files(context.Background(), params)
			if err != nil {
				return nil, errorsmod.Wrap(err, "Can't find file")
			}

			// Not 100% sure it's going to turn out right on the first go
			resForWasmBinding := bindings.FilesResponse{
				Files: queryRes.Files.String(),
			}

			bz, err := json.Marshal(resForWasmBinding)
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
