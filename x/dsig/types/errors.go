package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dsig module sentinel errors
var (
	NoFile        = sdkerrors.Register(ModuleName, 1300, "")
	NotOwner      = sdkerrors.Register(ModuleName, 1400, "")
	InvalidSignee = sdkerrors.Register(ModuleName, 1500, "")
	NoForm        = sdkerrors.Register(ModuleName, 1600, "")
	BadVote       = sdkerrors.Register(ModuleName, 1700, "")
	BadUser       = sdkerrors.Register(ModuleName, 1800, "")
	DuplicateFid  = sdkerrors.Register(ModuleName, 1900, "")
	DuplicateForm = sdkerrors.Register(ModuleName, 2000, "")
)
