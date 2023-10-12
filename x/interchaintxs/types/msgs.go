package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = (*MsgRegisterInterchainAccount)(nil) // _ sdk.HasValidateBasic = (*MsgRegisterInterchainAccount)(nil)
// our version of the sdk doesn't have the above function

// NewMsgRegisterInterchainAccount creates a new instance of MsgRegisterInterchainAccount
func NewMsgRegisterInterchainAccount(fromAddress, connectionID, interchainAccountID string) *MsgRegisterInterchainAccount {
	return &MsgRegisterInterchainAccount{
		FromAddress:         fromAddress,
		ConnectionId:        connectionID,
		InterchainAccountId: interchainAccountID,
	}
}
