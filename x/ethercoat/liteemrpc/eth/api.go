package eth

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"

	bech32 "github.com/btcsuite/btcutil/bech32"
	sdkbech32 "github.com/cosmos/cosmos-sdk/types/bech32"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/params"

	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdktx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/jackalLabs/canine-chain/x/ethercoat/liteemrpc/eth/backend"
	rpctypes "github.com/jackalLabs/canine-chain/x/ethercoat/liteemrpc/types"
	"github.com/jackalLabs/canine-chain/x/ethercoat/types"

	maintypes "github.com/jackalLabs/canine-chain/x/ethercoat/types"

	ethermint "github.com/jackalLabs/canine-chain/x/ethercoat/minttypes"
)

// PublicAPI is the eth_ prefixed set of APIs in the Web3 JSON-RPC spec.
type PublicAPI struct {
	ctx          context.Context
	clientCtx    client.Context
	queryClient  *rpctypes.QueryClient
	chainIDEpoch *big.Int
	logger       log.Logger
	backend      backend.EVMBackend
	nonceLock    *rpctypes.AddrLocker
	signer       ethtypes.Signer
}

// NewPublicAPI creates an instance of the public ETH Web3 API.
func NewPublicAPI(
	logger log.Logger,
	clientCtx client.Context,
	backend backend.EVMBackend,
	nonceLock *rpctypes.AddrLocker,
) *PublicAPI {
	eip155ChainID, err := ethermint.ParseChainID(clientCtx.ChainID)
	if err != nil {
		panic(err)
	}

	// overwriting the cosmos keyring with eth-compatible signatures
	// The signer used by the API should always be the 'latest' known one because we expect
	// signers to be backwards-compatible with old transactions.
	cfg := backend.ChainConfig()
	if cfg == nil {
		cfg = maintypes.DefaultChainConfig().EthereumConfig(eip155ChainID)
	}

	signer := ethtypes.LatestSigner(cfg)

	api := &PublicAPI{
		ctx:          context.Background(),
		clientCtx:    clientCtx,
		queryClient:  rpctypes.NewQueryClient(clientCtx),
		chainIDEpoch: eip155ChainID,
		logger:       logger.With("client", "json-rpc"),
		backend:      backend,
		nonceLock:    nonceLock,
		signer:       signer,
	}

	return api
}

// ClientCtx returns client context
func (e *PublicAPI) ClientCtx() client.Context {
	return e.clientCtx
}

func (e *PublicAPI) QueryClient() *rpctypes.QueryClient {
	return e.queryClient
}

func (e *PublicAPI) Ctx() context.Context {
	return e.ctx
}

// WORKING
// ProtocolVersion returns the supported Ethereum protocol version.
func (e *PublicAPI) ProtocolVersion() hexutil.Uint {
	e.logger.Error("eth_protocolVersion")
	return hexutil.Uint(65) // currently hardcoded as eth65
}

// WORKING
// ChainId is the EIP-155 replay-protection chain id for the current ethereum chain config.
func (e *PublicAPI) ChainId() (*hexutil.Big, error) { // nolint
	e.logger.Error("eth_chainId")
	// check if config.ChainID is used elsewhere
	return (*hexutil.Big)(e.chainIDEpoch), nil
}

// WORKING -- Always false.
// Mining returns whether or not this node is currently mining. Always false.
func (e *PublicAPI) Mining() bool {
	e.logger.Error("eth_mining")
	return false
}

// WORKING
// Hashrate returns the current node's hashrate. Always 0.
func (e *PublicAPI) Hashrate() hexutil.Uint64 {
	e.logger.Error("eth_hashrate")
	return 0
}

// WORKING
// GasPrice returns the current gas price based on min-gas-prices in config.yml
func (e *PublicAPI) GasPrice() (*hexutil.Big, error) {
	e.logger.Error("eth_gasPrice")
	// initalizing the result variables
	var result *big.Int
	var ujklprice float64

	// finding the working directory
	wd, _ := os.Getwd()
	// finding the canine-chain directory --
	// config.yml must be in canine-chain/config.yml
	i := strings.Index(wd, "canine-chain")
	if i < 0 {
		return (*hexutil.Big)(result), nil
	}
	cfgdir := wd[0:i]

	// importing config.yml
	fmt.Println(cfgdir + "config.yml")
	readcfg, err := os.ReadFile(cfgdir + "canine-chain/config.yml")
	if err != nil {
		e.logger.Error(fmt.Sprint(err))
		return (*hexutil.Big)(result), err
	}
	cfg := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal([]byte(readcfg), &cfg)
	if err2 != nil {
		e.logger.Error(fmt.Sprint(err))
		return (*hexutil.Big)(result), err2
	}
	// collecting the min-gas-prices string
	mingasprices := cfg["init"].(map[string]interface{})["app"].(map[string]interface{})["minimum-gas-prices"]
	if mingasprices == nil {
		e.logger.Error("ujkl min-gas-price either not configured or configured with different hierarchy")
	}
	// selecting ujkl only
	mgp_split := strings.Split(mingasprices.(string), ";")

	for _, price := range mgp_split {
		if i := strings.Index(price, "ujkl"); i != -1 {
			pricefloat, err := strconv.ParseFloat(price[0:i], 64)
			if err != nil {
				e.logger.Error("Failed to parse min-gas-prices into float")
				return (*hexutil.Big)(result), err
			}
			ujklprice = pricefloat
		}
	}
	// converting ujkl to GWEI (giga-wei)
	result = big.NewInt(int64(ujklprice * 10e8))
	return (*hexutil.Big)(result), nil
}

// WORKING
// Accounts returns the list of accounts available to this node.
func (e *PublicAPI) Accounts() ([]common.Address, error) {
	e.logger.Error("eth_accounts")

	addresses := make([]common.Address, 0) // return [] instead of nil if empty

	infos, err := e.clientCtx.Keyring.List()
	if err != nil {
		return addresses, err
	}

	for _, info := range infos {
		addressBytes := info.GetPubKey().Address().Bytes()
		e.logger.Error(common.BytesToAddress(addressBytes).String())
		addresses = append(addresses, common.BytesToAddress(addressBytes))
	}

	return addresses, nil
}

// WORKING
// BlockNumber returns the current block number.
func (e *PublicAPI) BlockNumber() (hexutil.Uint64, error) {
	e.logger.Error("eth_blockNumber")
	return e.backend.BlockNumber()
}

// WORKING
// GetBalance returns the provided account's balance up to the provided block number.
func (e *PublicAPI) GetBalance(address common.Address, blockNrOrHash rpctypes.BlockNumberOrHash) (*hexutil.Big, error) {
	e.logger.Error("eth_getBalance", "address", address.String(), "block number or hash", blockNrOrHash)
	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return nil, err
	}

	// converting to base32
	conv, err := bech32.ConvertBits(address.Bytes(), 8, 5, true)
	if err != nil {
		// implement better error catching
		fmt.Println("Error:", err)
	}
	encoded, err := bech32.Encode("jkl", conv)
	if err != nil {
		fmt.Println("Error:", err)
	}
	req := &banktypes.QueryBalanceRequest{
		Address: encoded,
		Denom:   "ujkl",
	}
	res, err := e.queryClient.BankQueryClient.Balance(rpctypes.ContextWithHeight(blockNum.Int64()), req)
	if err != nil {
		return nil, err
	}

	val, ok := sdk.NewIntFromString(res.Balance.Amount.String())
	if !ok {
		return nil, errors.New("invalid balance")
	}

	// rounding up
	rounding := sdk.NewInt(1000000000000)
	nval := val.Mul(rounding)
	e.logger.Error(rounding.String())
	return (*hexutil.Big)(nval.BigInt()), nil
}

// WORKING
// GetTransactionCount returns the number of transactions at the given address up to the given block number.
func (e *PublicAPI) GetTransactionCount(address common.Address, blockNrOrHash rpctypes.BlockNumberOrHash) (*hexutil.Uint64, error) {
	e.logger.Error("eth_getTransactionCount", "address", address.Hex(), "block number or hash", blockNrOrHash)
	blockNum, err := e.getBlockNumber(blockNrOrHash)
	if err != nil {
		return nil, err
	}
	return e.backend.GetTransactionCount(address, blockNum)
}

// GetBlockTransactionCountByHash returns the number of transactions in the block identified by hash.
func (e *PublicAPI) GetBlockTransactionCountByHash(hash common.Hash) *hexutil.Uint {
	e.logger.Error("eth_getBlockTransactionCountByHash", "hash", hash.Hex())

	block, err := e.clientCtx.Client.BlockByHash(e.ctx, hash.Bytes())
	if err != nil {
		e.logger.Debug("block not found", "hash", hash.Hex(), "error", err.Error())
		return nil
	}

	txs := hexutil.Uint(len(block.Block.Txs))
	return &txs
}

// / WORKING
// GetUncleCountByBlockHash returns the number of uncles in the block identified by hash. Always zero.
func (e *PublicAPI) GetUncleCountByBlockHash(hash common.Hash) hexutil.Uint {
	return 0
}

// WORKING
// GetUncleCountByBlockNumber returns the number of uncles in the block identified by number. Always zero.
func (e *PublicAPI) GetUncleCountByBlockNumber(blockNum rpctypes.BlockNumber) hexutil.Uint {
	return 0
}

// WORKING
// GetCode returns the contract code at the given address and block number.
func (e *PublicAPI) GetCode(address common.Address, blockNrOrHash rpctypes.BlockNumberOrHash) (hexutil.Bytes, error) {
	e.logger.Error("eth_getCode", "address", address.Hex(), "block number or hash", blockNrOrHash)

	// returning a hash of an empty string, as no EVM contracts are running on Canine.
	var ret hexutil.Bytes
	return ret, nil
}

// SendRawTransaction send a raw Ethereum transaction.
func (e *PublicAPI) SendRawTransaction(data hexutil.Bytes) (common.Hash, error) {
	e.logger.Error("eth_sendRawTransaction", "length", len(data))

	// RLP decode raw transaction bytes
	tx := &ethtypes.Transaction{}
	if err := tx.UnmarshalBinary(data); err != nil {
		e.logger.Error("transaction decoding failed", "error", err.Error())
		return common.Hash{}, err
	}

	// check the local node config in case unprotected txs are disabled
	if !tx.Protected() {
		// Ensure only eip155 signed transactions are submitted if EIP155Required is set.
		return common.Hash{}, errors.New("only replay-protected (EIP-155) transactions allowed over RPC")
	}

	ethereumTx := &types.MsgEthereumTx{}
	if err := ethereumTx.FromEthereumTx(tx); err != nil {
		e.logger.Error("transaction converting failed", "error", err.Error())
		return common.Hash{}, err
	}

	if err := ethereumTx.ValidateBasic(); err != nil {
		e.logger.Debug("tx failed basic validation", "error", err.Error())
		return common.Hash{}, err
	}

	cosmosTx, err := ethereumTx.BuildTx(e.clientCtx.TxConfig.NewTxBuilder(), "ujkl")
	if err != nil {
		e.logger.Error("failed to build cosmos tx", "error", err.Error())
		return common.Hash{}, err
	}

	// decrypting a sample cosmos transaction
	rawtxbytes := "Co8BCooBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEmoKKmprbDFmaHduaHJzOHN2dzJzYzdjdHlrZ3AzeWthbDdxc2owcjM0a25zaBIqamtsMXZ2anV6MjAwODZ1amp3azNybW5maHk2NzdseG43aHJ2dTR6emZyGhAKBHVqa2wSCDEyMDAwMDAwEgASZgpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAisET0CjmTWbbo6/Gk0Eb2AcJWtwcjdgMmqG7H9BOqTwEgQKAgh/GAESEgoMCgR1amtsEgQxOTg1EKLsBBpAeX2AqFLHLrasQ74+HLkuo45hurKT+TH/BST5p3P6xksm5v1Va/oGgUr6c4xrdNkeexbfbnUeQ4uYEK438s/L3w=="

	// copied from the default tx processor found in x/auth
	rawDecodeTx, err := base64.StdEncoding.DecodeString(rawtxbytes)
	if err != nil {
		e.logger.Error(fmt.Sprint(err))
	}

	txstruct, err := e.clientCtx.TxConfig.TxDecoder()(rawDecodeTx)
	if err != nil {
		e.logger.Error(fmt.Sprint(err))
	}

	jsontx, err := e.clientCtx.TxConfig.TxJSONEncoder()(txstruct)
	if err != nil {
		e.logger.Error(fmt.Sprint(err))
	}

	e.clientCtx.PrintBytes(jsontx)

	// Encode transaction by default Tx encoder
	txBytes, err := e.clientCtx.TxConfig.TxEncoder()(cosmosTx)
	if err != nil {
		e.logger.Error("failed to encode eth tx using default encoder", "error", err.Error())
		return common.Hash{}, err
	}

	txHash := ethereumTx.AsTransaction().Hash()

	syncCtx := e.clientCtx.WithBroadcastMode(flags.BroadcastSync)
	rsp, err := syncCtx.BroadcastTx(txBytes)
	if rsp != nil && rsp.Code != 0 {
		err = sdkerrors.ABCIError(rsp.Codespace, rsp.Code, rsp.RawLog)
	}
	if err != nil {
		e.logger.Error("failed to broadcast tx", "error", err.Error())
		return txHash, err
	}

	return txHash, nil
}

// checkTxFee is an internal function used to check whether the fee of
// the given transaction is _reasonable_(under the cap).
func checkTxFee(gasPrice *big.Int, gas uint64, cap float64) error {
	// Short circuit if there is no cap for transaction fee at all.
	if cap == 0 {
		return nil
	}
	totalfee := new(big.Float).SetInt(new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gas)))
	// 1 photon in 10^18 aphoton
	oneToken := new(big.Float).SetInt(big.NewInt(params.Ether))
	// quo = rounded(x/y)
	feeEth := new(big.Float).Quo(totalfee, oneToken)
	// no need to check error from parsing
	feeFloat, _ := feeEth.Float64()
	if feeFloat > cap {
		return fmt.Errorf("tx fee (%.2f ether) exceeds the configured cap (%.2f ether)", feeFloat, cap)
	}
	return nil
}

// WORKING
// EstimateGas returns an estimate of gas usage for the given smart contract call.
func (e *PublicAPI) EstimateGas(args maintypes.TransactionArgs, blockNrOptional *rpctypes.BlockNumber) (hexutil.Uint64, error) {
	e.logger.Error("eth_estimateGas")
	//  checking if empty transaction
	if args.To == nil {
		return hexutil.Uint64(0), fmt.Errorf("No reciever address")
	}
	// converting addresses
	// converting 'from' address
	fromAddHex := args.From
	fromBech32, err := sdkbech32.ConvertAndEncode("jkl", fromAddHex.Bytes())
	if err != nil {
		return hexutil.Uint64(0), fmt.Errorf("encoding bech32 failed: %w", err)
	}
	// converting 'to' address
	toAddHex := args.To
	toBech32, err := sdkbech32.ConvertAndEncode("jkl", toAddHex.Bytes())
	if err != nil {
		return hexutil.Uint64(0), fmt.Errorf("encoding bech32 failed: %w", err)
	}
	// transfer no coins
	var blank int64
	blank = 10
	sdkBlank := sdk.NewInt(blank)

	gasCoin := sdk.NewCoin("ujkl", sdkBlank)
	gasCoins := sdk.NewCoins(gasCoin)

	// converting the ethereum transaction into a sdk.Msg
	banktx := banktypes.MsgSend{
		FromAddress: fromBech32,
		ToAddress:   toBech32,
		Amount:      gasCoins,
	}
	// creating blank fees and gas to send in mock transaction
	blankfees := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(1)))        // guessing 1
	blankcoins := sdk.NewDecCoins(sdk.NewDecCoin("ujkl", sdk.NewInt(0))) // setting to zero for estimate

	// calculating the blocknumber
	currentHeightHex, err := e.backend.BlockNumber()
	if err != nil {
		return hexutil.Uint64(0), err
	}
	currentHeightNum, err := hexutil.DecodeUint64(currentHeightHex.String())
	if err != nil {
		return hexutil.Uint64(0), err
	}

	// creating a new txfactory
	var txf sdktx.Factory
	txf = txf.WithTxConfig(e.clientCtx.TxConfig).
		WithKeybase(e.clientCtx.Keyring).
		WithAccountRetriever(e.clientCtx.AccountRetriever).
		WithGas(gasCoin.Amount.Abs().Uint64()).
		WithTimeoutHeight(currentHeightNum + 100).
		WithChainID(e.clientCtx.ChainID).
		WithFees(blankfees.String()).
		WithGasPrices(blankcoins.String()).
		WithSimulateAndExecute(true).
		WithGasAdjustment(1.0)

	// sending the sdk.Msg args
	_, gas, err := sdktx.CalculateGas(e.clientCtx, txf, &banktx)
	e.logger.Error(fmt.Sprint(gas), err)
	if err != nil {
		e.logger.Error("failed to calculate gas")
		return hexutil.Uint64(0), err
	}
	return hexutil.Uint64(gas), nil
}

// WORKING
// GetBlockByNumber returns the block identified by number.
func (e *PublicAPI) GetBlockByNumber(ethBlockNum rpctypes.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	e.logger.Error("eth_getBlockByNumber", "number", ethBlockNum, "full", fullTx)
	return e.backend.GetBlockByNumber(ethBlockNum, fullTx)
}

// GetUncleByBlockHashAndIndex returns the uncle identified by hash and index. Always returns nil.
func (e *PublicAPI) GetUncleByBlockHashAndIndex(hash common.Hash, idx hexutil.Uint) map[string]interface{} {
	return nil
}

// GetUncleByBlockNumberAndIndex returns the uncle identified by number and index. Always returns nil.
func (e *PublicAPI) GetUncleByBlockNumberAndIndex(number, idx hexutil.Uint) map[string]interface{} {
	return nil
}

// getBlockNumber returns the BlockNumber from BlockNumberOrHash
func (e *PublicAPI) getBlockNumber(blockNrOrHash rpctypes.BlockNumberOrHash) (rpctypes.BlockNumber, error) {
	switch {
	case blockNrOrHash.BlockHash == nil && blockNrOrHash.BlockNumber == nil:
		return rpctypes.EthEarliestBlockNumber, fmt.Errorf("types BlockHash and BlockNumber cannot be both nil")
	case blockNrOrHash.BlockHash != nil:
		blockNumber, err := e.backend.GetBlockNumberByHash(*blockNrOrHash.BlockHash)
		if err != nil {
			return rpctypes.EthEarliestBlockNumber, err
		}
		return rpctypes.NewBlockNumber(blockNumber), nil
	case blockNrOrHash.BlockNumber != nil:
		return *blockNrOrHash.BlockNumber, nil
	default:
		return rpctypes.EthEarliestBlockNumber, nil
	}
}
