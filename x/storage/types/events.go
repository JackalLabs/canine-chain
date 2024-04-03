package types

// storage module event types
const (
	EventTypeBuyStorage     = "buy_storage"
	EventTypeSignContract   = "sign_contract"
	EventTypeUpgradeStorage = "upgrade_storage"
	EventTypeCancelContract = "cancel_contract"

	AttributeValueCategory = ModuleName

	AttributeKeyBuyer       = "buyer" // buy storage
	AttributeKeyReceiver    = "receiver"
	AttributeKeyBytesBought = "bytes_bought"
	AttributeKeyTimeBought  = "hours_bought"

	AttributeKeySigner   = "signer" // sign storage deal
	AttributeKeyContract = "contract"
	AttributeKeyPayOnce  = "pay_once"
)
