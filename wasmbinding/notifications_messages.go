package wasmbinding

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	notificationskeeper "github.com/jackalLabs/canine-chain/v5/x/notifications/keeper"
	notificationstypes "github.com/jackalLabs/canine-chain/v5/x/notifications/types"
)

func PerformCreateNotification(s *notificationskeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, createNotification *notificationstypes.MsgCreateNotification) error {
	if createNotification == nil {
		return wasmvmtypes.InvalidRequest{Err: "create notification null error"}
	}

	if createNotification.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := createNotification.ValidateBasic(); err != nil {
		return err
	}

	msgServer := notificationskeeper.NewMsgServerImpl(*s)
	_, err := msgServer.CreateNotification(sdk.WrapSDKContext(ctx), createNotification)
	if err != nil {
		return sdkerrors.Wrap(err, "create notification error from message")
	}

	return nil
}

func PerformDeleteNotification(s *notificationskeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, deleteNotification *notificationstypes.MsgDeleteNotification) error {
	if deleteNotification == nil {
		return wasmvmtypes.InvalidRequest{Err: "delete notification null error"}
	}

	if deleteNotification.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := deleteNotification.ValidateBasic(); err != nil {
		return err
	}

	msgServer := notificationskeeper.NewMsgServerImpl(*s)
	_, err := msgServer.DeleteNotification(sdk.WrapSDKContext(ctx), deleteNotification)
	if err != nil {
		return sdkerrors.Wrap(err, "delete notification error from message")
	}

	return nil
}

func PerformBlockSenders(s *notificationskeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, blockSenders *notificationstypes.MsgBlockSenders) error {
	if blockSenders == nil {
		return wasmvmtypes.InvalidRequest{Err: "block senders null error"}
	}

	if blockSenders.Creator != contractAddr.String() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator of bindings is not bindings contract address")
	}

	if err := blockSenders.ValidateBasic(); err != nil {
		return err
	}

	msgServer := notificationskeeper.NewMsgServerImpl(*s)
	_, err := msgServer.BlockSenders(sdk.WrapSDKContext(ctx), blockSenders)
	if err != nil {
		return sdkerrors.Wrap(err, "block senders error from message")
	}

	return nil
}
