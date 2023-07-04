package wasmbinding

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/wasmbinding/bindings"
	filetreekeeper "github.com/jackalLabs/canine-chain/v3/x/filetree/keeper"
	filetreetypes "github.com/jackalLabs/canine-chain/v3/x/filetree/types"
	storagekeeper "github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
	storagetypes "github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (m *CustomMessenger) postAndSign(ctx sdk.Context, contractAddr sdk.AccAddress, postAndSign *bindings.PostAndSign, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformPostAndSign(m.filetree, m.storage, ctx, contractAddr, postAndSign, sender)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform buy storage")
	}
	return nil, nil, nil
}

// Do we care that this function is public? everything is gated by our fork of wasmd anyways
func PerformPostAndSign(f *filetreekeeper.Keeper, s *storagekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, postAndSign *bindings.PostAndSign, sender string) error {
	if postAndSign == nil {
		return wasmvmtypes.InvalidRequest{Err: "postAndSign msg is empty"}
	}

	filetreeMsgServer := filetreekeeper.NewMsgServerImpl(*f)

	msgPostFiles := filetreetypes.NewMsgPostFile(
		sender,
		postAndSign.Account,
		postAndSign.HashParent,
		postAndSign.HashChild,
		postAndSign.Contents,
		postAndSign.Viewers,
		postAndSign.Editors,
		postAndSign.TrackingNumber,
	)

	if err := msgPostFiles.ValidateBasic(); err != nil {
		return errorsmod.Wrap(err, "failed validating MsgPostFiles")
	}

	// Post files
	_, err := filetreeMsgServer.PostFile(
		sdk.WrapSDKContext(ctx),
		msgPostFiles,
	)
	if err != nil {
		return errorsmod.Wrap(err, "failed to post file:")
	}

	storageMsgServer := storagekeeper.NewMsgServerImpl(*s)

	msgSignContract := storagetypes.NewMsgSignContract(
		sender,
		postAndSign.Cid,
		postAndSign.PayOnce,
	)

	if err := msgSignContract.ValidateBasic(); err != nil {
		return errorsmod.Wrap(err, "failed validating msgSignContract")
	}

	// sign contract
	_, error := storageMsgServer.SignContract(
		sdk.WrapSDKContext(ctx),
		msgSignContract,
	)
	if error != nil {
		return errorsmod.Wrap(err, "failed to sign contract:")
	}
	return nil
}

func (m *CustomMessenger) deleteAndCancel(ctx sdk.Context, contractAddr sdk.AccAddress, deleteAndCancel *bindings.DeleteAndCancel, sender string) ([]sdk.Event, [][]byte, error) {
	err := PerformDeleteAndCancel(m.filetree, m.storage, ctx, contractAddr, deleteAndCancel, sender)
	if err != nil {
		return nil, nil, errorsmod.Wrap(err, "perform buy storage")
	}
	return nil, nil, nil
}

// Do we care that this function is public? everything is gated by our fork of wasmd anyways
func PerformDeleteAndCancel(f *filetreekeeper.Keeper, s *storagekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, deleteAndCancel *bindings.DeleteAndCancel, sender string) error {
	if deleteAndCancel == nil {
		return wasmvmtypes.InvalidRequest{Err: "deleteAndCancel is null"}
	}
	storageMsgServer := storagekeeper.NewMsgServerImpl(*s)

	var toRemove []string
	err := json.Unmarshal([]byte(deleteAndCancel.Cids), &toRemove)
	if err != nil {
		return errorsmod.Wrap(err, "failed to unmarshal Cids from 'deleteAndCancel'")
	}

	for _, cid := range toRemove {
		msgCancelContract := storagetypes.NewMsgCancelContract(
			sender,
			cid,
		)

		if err := msgCancelContract.ValidateBasic(); err != nil {
			return errorsmod.Wrap(err, "failed validating msgCancelContract")
		}

		// sign contract
		_, error := storageMsgServer.CancelContract(
			sdk.WrapSDKContext(ctx),
			msgCancelContract,
		)
		if error != nil {
			return errorsmod.Wrap(error, "failed to sign contract:")
		}
	}

	//=======================

	msgServer := filetreekeeper.NewMsgServerImpl(*f)

	msgDeleteFile := filetreetypes.NewMsgDeleteFile(
		sender,
		deleteAndCancel.HashPath,
		deleteAndCancel.Account,
	)

	if err := msgDeleteFile.ValidateBasic(); err != nil {
		return errorsmod.Wrap(err, "failed validating msgDeleteFile")
	}

	// Delete File
	_, error := msgServer.DeleteFile(
		sdk.WrapSDKContext(ctx),
		msgDeleteFile,
	)
	if error != nil {
		return errorsmod.Wrap(err, "failed to delete file:")
	}
	return nil
}
