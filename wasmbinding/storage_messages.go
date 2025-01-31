package wasmbinding

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	storagekeeper "github.com/jackalLabs/canine-chain/v4/x/storage/keeper"
	storagetypes "github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

func PerformPostFile(s *storagekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, postFile *storagetypes.MsgPostFile) error {
	if postFile == nil {
		return wasmvmtypes.InvalidRequest{Err: "post file null error"}
	}

	if postFile.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := postFile.ValidateBasic(); err != nil {
		return err
	}

	msgServer := storagekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.PostFile(sdk.WrapSDKContext(ctx), postFile)
	if err != nil {
		return sdkerrors.Wrap(err, "post file error from message")
	}

	return nil
}

func PerformDeleteFile(s *storagekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, deleteFile *storagetypes.MsgDeleteFile) error {
	if deleteFile == nil {
		return wasmvmtypes.InvalidRequest{Err: "delete file null error"}
	}

	if deleteFile.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := deleteFile.ValidateBasic(); err != nil {
		return err
	}

	msgServer := storagekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.DeleteFile(sdk.WrapSDKContext(ctx), deleteFile)
	if err != nil {
		return sdkerrors.Wrap(err, "delete file error from message")
	}

	return nil
}

func PerformBuyStorage(s *storagekeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, buyStorage *storagetypes.MsgBuyStorage) error {
	if buyStorage == nil {
		return wasmvmtypes.InvalidRequest{Err: "buy storage null error"}
	}

	if buyStorage.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := buyStorage.ValidateBasic(); err != nil {
		return err
	}

	msgServer := storagekeeper.NewMsgServerImpl(*s)
	_, err := msgServer.BuyStorage(sdk.WrapSDKContext(ctx), buyStorage)
	if err != nil {
		return sdkerrors.Wrap(err, "buy storage error from message")
	}

	return nil
}
