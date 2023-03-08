package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/notifications module sentinel errors
var (
	ErrNotiCounterNotFound    = sdkerrors.Register(ModuleName, 1100, "User's notiCounter not set")
	ErrCantUnmarshall         = sdkerrors.Register(ModuleName, 1101, "Cannot unmarshall from JSON")
	ErrBlockedSender          = sdkerrors.Register(ModuleName, 1102, "You are a blocked sender")
	ErrNotificationAlreadySet = sdkerrors.Register(ModuleName, 1103, "Notification already set")
	ErrOnlyOwnerCanBlock      = sdkerrors.Register(ModuleName, 1104, "Only the notiCounter owner can block a sender")
)
