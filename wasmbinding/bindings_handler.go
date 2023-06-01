package wasmbinding

import (
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// WARNING
// Any security vulnerabilities with declaring these types at all?

// I don't think this encoder is going to work, I think I need to implement the method some more

// msgEncoder is an extension point to customize encodings
type msgEncoder interface {
	// Encode converts wasmvm message to n cosmos message types
	Encode(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Msg, error)
}

// MessageRouter ADR 031 request type routing
type MessageRouter interface {
	Handler(msg sdk.Msg) baseapp.MsgServiceHandler
}

// SDKMessageHandler can handles messages that can be encoded into sdk.Message types and routed.
type SDKMessageHandler struct {
	router   MessageRouter
	encoders msgEncoder
}
