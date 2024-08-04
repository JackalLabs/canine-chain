package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/notifications module sentinel errors
var (
	ErrCantUnmarshall         = sdkerrors.Register(ModuleName, 1101, "cannot unmarshall from JSON")
	ErrBlockedSender          = sdkerrors.Register(ModuleName, 1102, "you are a blocked sender")
	ErrNotificationAlreadySet = sdkerrors.Register(ModuleName, 1103, "notification already set")
	ErrNotificationNotFound   = sdkerrors.Register(ModuleName, 1105, "notification does not exist")
	ErrNotNotificationOwner   = sdkerrors.Register(ModuleName, 1106, "you do not own this notification")

	ErrInvalidContents = sdkerrors.Register(ModuleName, 1110, "contents must be valid JSON")
)
