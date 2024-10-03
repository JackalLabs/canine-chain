package types

const (
	EventSetPrimaryName = "set_primary_name"
	EventSetName        = "set_name"
	EventRemoveName     = "remove_name"
	EventSetBid         = "add_bid"
	EventAcceptBid      = "accept_bid"
	EventAddRecord      = "add_record"
	EventBuyName        = "buy_name"
	EventRemoveBid      = "remove_bid"
	EventRemoveRecord   = "remove_record"
	EventRemoveSale     = "remove_sale"
	EventSetSale        = "add_sale"
	EventInit           = "init_rns"
	EventRegister       = "register"
	EventTransfer       = "transfer_name"
	EventUpdate         = "update"

	AttributeName     = "name"
	AttributeOwner    = "owner"
	AttributeValue    = "value"
	AttributeExpires  = "expires"
	AttributeBidder   = "bidder"
	AttributePrice    = "price"
	AttributeReceiver = "receiver"

	AttributeValueCategory = ModuleName

	AttributeKeySigner = "signer" // sign storage deal

	EventTypeJackalMessage = "jackal_message"
)
