package wasmbinding

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	filetreekeeper "github.com/jackalLabs/canine-chain/v5/x/filetree/keeper"
	filetreetypes "github.com/jackalLabs/canine-chain/v5/x/filetree/types"
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

func PerformAddViewers(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, addViewers *filetreetypes.MsgAddViewers) error {
	if addViewers == nil {
		return wasmvmtypes.InvalidRequest{Err: "add viewers null error"}
	}

	if addViewers.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := addViewers.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.AddViewers(sdk.WrapSDKContext(ctx), addViewers)
	if err != nil {
		return sdkerrors.Wrap(err, "add viewers error from message")
	}

	return nil
}

func PerformPostKey(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, postKey *filetreetypes.MsgPostKey) error {
	if postKey == nil {
		return wasmvmtypes.InvalidRequest{Err: "post key null error"}
	}

	if postKey.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := postKey.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.PostKey(sdk.WrapSDKContext(ctx), postKey)
	if err != nil {
		return sdkerrors.Wrap(err, "post key error from message")
	}

	return nil
}

func PerformDeleteFileTree(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, deleteFileTree *filetreetypes.MsgDeleteFile) error {
	if deleteFileTree == nil {
		return wasmvmtypes.InvalidRequest{Err: "delete file tree null error"}
	}

	if deleteFileTree.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := deleteFileTree.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.DeleteFile(sdk.WrapSDKContext(ctx), deleteFileTree)
	if err != nil {
		return sdkerrors.Wrap(err, "delete file tree error from message")
	}

	return nil
}

func PerformRemoveViewers(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, removeViewers *filetreetypes.MsgRemoveViewers) error {
	if removeViewers == nil {
		return wasmvmtypes.InvalidRequest{Err: "remove viewers null error"}
	}

	if removeViewers.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := removeViewers.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.RemoveViewers(sdk.WrapSDKContext(ctx), removeViewers)
	if err != nil {
		return sdkerrors.Wrap(err, "remove viewers error from message")
	}

	return nil
}

func PerformProvisionFileTree(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, provisionFileTree *filetreetypes.MsgProvisionFileTree) error {
	if provisionFileTree == nil {
		return wasmvmtypes.InvalidRequest{Err: "provision file tree null error"}
	}

	if provisionFileTree.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := provisionFileTree.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.ProvisionFileTree(sdk.WrapSDKContext(ctx), provisionFileTree)
	if err != nil {
		return sdkerrors.Wrap(err, "provision file tree error from message")
	}

	return nil
}

func PerformAddEditors(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, addEditors *filetreetypes.MsgAddEditors) error {
	if addEditors == nil {
		return wasmvmtypes.InvalidRequest{Err: "add editors null error"}
	}

	if addEditors.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := addEditors.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.AddEditors(sdk.WrapSDKContext(ctx), addEditors)
	if err != nil {
		return sdkerrors.Wrap(err, "add editors error from message")
	}

	return nil
}

func PerformRemoveEditors(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, removeEditors *filetreetypes.MsgRemoveEditors) error {
	if removeEditors == nil {
		return wasmvmtypes.InvalidRequest{Err: "remove editors null error"}
	}

	if removeEditors.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := removeEditors.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.RemoveEditors(sdk.WrapSDKContext(ctx), removeEditors)
	if err != nil {
		return sdkerrors.Wrap(err, "remove editors error from message")
	}

	return nil
}

func PerformResetEditors(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, resetEditors *filetreetypes.MsgResetEditors) error {
	if resetEditors == nil {
		return wasmvmtypes.InvalidRequest{Err: "reset editors null error"}
	}

	if resetEditors.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := resetEditors.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.ResetEditors(sdk.WrapSDKContext(ctx), resetEditors)
	if err != nil {
		return sdkerrors.Wrap(err, "reset editors error from message")
	}

	return nil
}

func PerformResetViewers(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, resetViewers *filetreetypes.MsgResetViewers) error {
	if resetViewers == nil {
		return wasmvmtypes.InvalidRequest{Err: "reset viewers null error"}
	}

	if resetViewers.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := resetViewers.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.ResetViewers(sdk.WrapSDKContext(ctx), resetViewers)
	if err != nil {
		return sdkerrors.Wrap(err, "reset viewers error from message")
	}

	return nil
}

func PerformChangeOwner(s *filetreekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, changeOwner *filetreetypes.MsgChangeOwner) error {
	if changeOwner == nil {
		return wasmvmtypes.InvalidRequest{Err: "change owner null error"}
	}

	if changeOwner.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := changeOwner.ValidateBasic(); err != nil {
		return err
	}

	msgServer := filetreekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.ChangeOwner(sdk.WrapSDKContext(ctx), changeOwner)
	if err != nil {
		return sdkerrors.Wrap(err, "change owner error from message")
	}

	return nil
}
