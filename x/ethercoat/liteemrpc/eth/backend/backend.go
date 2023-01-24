package backend

import (
	"context"
	"math/big"

	"github.com/cosmos/cosmos-sdk/server/config"

	emrpctypes "github.com/jackalLabs/canine-chain/x/ethercoat/liteemrpc/types"

	ethermint "github.com/jackalLabs/canine-chain/x/ethercoat/minttypes"

	"github.com/cosmos/cosmos-sdk/client"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	maintypes "github.com/jackalLabs/canine-chain/x/ethercoat/types"

	servertypes "github.com/cosmos/cosmos-sdk/server"

	"github.com/tendermint/tendermint/libs/log"
	tmrpctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// BackendI implements the Cosmos and EVM backend.
type BackendI interface { // nolint: revive
	CosmosBackend
	EVMBackend
}

// CosmosBackend implements the functionality shared within cosmos namespaces
// as defined by Wallet Connect V2: https://docs.walletconnect.com/2.0/json-rpc/cosmos.
// Implemented by Backend.
type CosmosBackend interface { // TODO: define
	// GetAccounts()
	// SignDirect()
	// SignAmino()
}

// EVMBackend implements the functionality shared within ethereum namespaces
// as defined by EIP-1474: https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1474.md
// Implemented by Backend.
type EVMBackend interface {
	// To keep
	BlockNumber() (hexutil.Uint64, error)
	GetTransactionCount(address common.Address, blockNum emrpctypes.BlockNumber) (*hexutil.Uint64, error)
	ChainConfig() *params.ChainConfig
	GetBlockByNumber(blockNum emrpctypes.BlockNumber, fullTx bool) (map[string]interface{}, error)
	GetTendermintBlockByNumber(blockNum emrpctypes.BlockNumber) (*tmrpctypes.ResultBlock, error)
	GetTendermintBlockResultByNumber(height *int64) (*tmrpctypes.ResultBlockResults, error)
	GetEthereumMsgsFromTendermintBlock(block *tmrpctypes.ResultBlock, blockRes *tmrpctypes.ResultBlockResults) []*maintypes.MsgEthereumTx
	BaseFee(blockRes *tmrpctypes.ResultBlockResults) (*big.Int, error)
	GetBlockNumberByHash(blockHash common.Hash) (*big.Int, error)
	GetTendermintBlockByHash(blockHash common.Hash) (*tmrpctypes.ResultBlock, error)
	BlockByNumber(blockNum emrpctypes.BlockNumber) (*ethtypes.Block, error)
	BlockByHash(blockHash common.Hash) (*ethtypes.Block, error)
	GetBlockByHash(hash common.Hash, fullTx bool) (map[string]interface{}, error)
}

var _ BackendI = (*Backend)(nil)

// Backend implements the BackendI interface
type Backend struct {
	ctx                 context.Context
	clientCtx           client.Context
	queryClient         *emrpctypes.QueryClient // gRPC query client
	authQueryClient     authtypes.QueryClient
	bankQueryClient     banktypes.QueryClient
	logger              log.Logger
	chainID             *big.Int
	cfg                 config.Config
	allowUnprotectedTxs bool
}

// NewBackend creates a new Backend instance for cosmos and ethereum namespaces
func NewBackend(ctx *servertypes.Context, logger log.Logger, clientCtx client.Context, allowUnprotectedTxs bool) *Backend {
	chainID, err := ethermint.ParseChainID(clientCtx.ChainID)
	if err != nil {
		panic(err)
	}

	appConf, _ := config.GetConfig(ctx.Viper)
	return &Backend{
		ctx:                 context.Background(),
		clientCtx:           clientCtx,
		queryClient:         emrpctypes.NewQueryClient(clientCtx),
		authQueryClient:     authtypes.NewQueryClient(clientCtx),
		logger:              logger.With("module", "backend"),
		chainID:             chainID,
		cfg:                 appConf,
		allowUnprotectedTxs: allowUnprotectedTxs,
	}
}
