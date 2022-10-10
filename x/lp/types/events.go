package types

const (
	TypedEventPoolInfo    = "pool_info"
	TypedEventPoolCreated = "pool_created"
	TypedEventPoolJoined  = "pool_joined"
	TypedEventPoolExited  = "pool_exited"
	TypedEventCoinSwapped = "pool_swapped"

	AttrValueModule = ModuleName

	AttrKeyPoolId       = "pool_id"
	AttrKeyPoolBalance  = "pool_balance"
	AttrKeySwapFee      = "swap_fee"
	AttrKeyPenaltyFee   = "penalty_fee"
	AttrKeyCoinsIn      = "coins_in"
	AttrKeyCoinsOut     = "coins_out"
	AttrKeySwapFeeMulti = "swap fee multiplier"
	AttrKeyPenaltyMulti = "penalty multiplier"
	AttrKeyLPTokenDenom = "liquidity pool token denom"
	AttrKeyLockDuration = "lock duration"
)
