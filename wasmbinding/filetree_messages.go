package wasmbinding

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	filetreekeeper "github.com/jackalLabs/canine-chain/v4/x/filetree/keeper"
	filetreetypes "github.com/jackalLabs/canine-chain/v4/x/filetree/types"
)

func PerformPostFileTree(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, postFileTree *filetreetypes.MsgPostFile) error {
	if postFileTree == nil {
		return wasmvmtypes.InvalidRequest{Err: "post file tree null error"}
	}

	if postFileTree.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := postFileTree.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.PostFile(sdk.WrapSDKContext(ctx), postFileTree)
	if err != nil {
		return sdkerrors.Wrap(err, "post file tree error from message")
	}

	return nil
}
