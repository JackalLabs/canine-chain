package types

const (
	TypedEventPoolInfo       = "pool_info"
	TypedEventPoolCreated    = "pool_created"
	TypedEventPoolJoined     = "pool_joined"
	TypedEventPoolExited     = "pool_exited"
	TypedEventCoinSwapped    = "pool_swapped"
	TypedEventCoinSwapFailed = "pool_swap_failed"

	AttrValueModule = ModuleName

	AttrKeyPoolId       = "pool_id"
	AttrKeyPoolBalance  = "pool_balance"
	AttrKeySwapFee      = "swap_fee"
	AttrKeyProtocolFee  = "protocol_fee"
	AttrKeyPenaltyFee   = "penalty_fee"
	AttrKeyCoinsIn      = "coins_in"
	AttrKeyCoinsOut     = "coins_out"
	AttrKeyMinCoinsOut  = "minimum_coins_out"
	AttrKeySwapFeeMulti = "swap_fee_multiplier"
	AttrKeyPenaltyMulti = "penalty_multiplier"
	AttrKeyLPTokenDenom = "liquidity_pool_token_denom"
	AttrKeyLockDuration = "lock_duration"
)
