package types

// storage module event types
const (
	EventTypeBuyStorage     = "buy_storage"
	EventTypeSignContract   = "post_file"
	EventTypeCancelContract = "delete_file"

	AttributeValueCategory = ModuleName

	AttributeKeyBuyer       = "buyer" // buy storage
	AttributeKeyReceiver    = "receiver"
	AttributeKeyBytesBought = "bytes_bought"
	AttributeKeyTimeBought  = "hours_bought"

	AttributeKeySigner   = "signer" // sign storage deal
	AttributeKeyContract = "file"
	AttributeKeyPayOnce  = "pay_once"
	AttributeKeyStart    = "start"

	EventTypeJackalMessage = "jackal_message"
)
