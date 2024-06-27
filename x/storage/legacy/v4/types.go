package v4

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/params/types"
	types2 "github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

type LegacyParams struct {
	DepositAccount string `protobuf:"bytes,1,opt,name=deposit_account,json=depositAccount,proto3" json:"deposit_account,omitempty"`
	ProofWindow    int64  `protobuf:"varint,2,opt,name=proof_window,json=proofWindow,proto3" json:"proof_window,omitempty"`
	// Chunk size of a file is divided into
	// The value cannot be smaller than 1 to avoid zero division
	ChunkSize    int64  `protobuf:"varint,3,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	MissesToBurn int64  `protobuf:"varint,4,opt,name=misses_to_burn,json=missesToBurn,proto3" json:"misses_to_burn,omitempty"`
	PriceFeed    string `protobuf:"bytes,5,opt,name=price_feed,json=priceFeed,proto3" json:"price_feed,omitempty"`
	// Life span of a contract in blocks
	MaxContractAgeInBlocks int64 `protobuf:"varint,6,opt,name=max_contract_age_in_blocks,json=maxContractAgeInBlocks,proto3" json:"max_contract_age_in_blocks,omitempty"`
	PricePerTbPerMonth     int64 `protobuf:"varint,7,opt,name=price_per_tb_per_month,json=pricePerTbPerMonth,proto3" json:"price_per_tb_per_month,omitempty"`
}

// ParamSetPairs get the params.ParamSet
func (p *LegacyParams) ParamSetPairs() types.ParamSetPairs {
	return types.ParamSetPairs{
		types.NewParamSetPair(types2.KeyDepositAccount, &p.DepositAccount, validateDeposit),
		types.NewParamSetPair(types2.KeyProofWindow, &p.ProofWindow, validateProofWindow),
		types.NewParamSetPair(types2.KeyChunkSize, &p.ChunkSize, validateChunkSize),
		types.NewParamSetPair(types2.KeyMissesToBurn, &p.MissesToBurn, validateMissesToBurn),
		types.NewParamSetPair(types2.KeyPriceFeed, &p.PriceFeed, validatePriceFeed),
		types.NewParamSetPair(
			types2.KeyMaxContractAgeInBlocks,
			&p.MaxContractAgeInBlocks,
			validateMaxContractAgeInBlocks),
		types.NewParamSetPair(
			types2.KeyPricePerTbPerMonth,
			&p.PricePerTbPerMonth,
			validatePricePerTbPerMonth),
	}
}

func validateProofWindow(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 1 {
		return errors.New("proof window must be greater than 1")
	}

	return nil
}

func validateDeposit(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("deposit cannot be blank")
	}

	return nil
}

func validateChunkSize(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 1 {
		return errors.New("chunk size cannot be smaller than 1")
	}

	return nil
}

func validateMissesToBurn(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 1 {
		return errors.New("misses to burn cannot be smaller than 1")
	}

	return nil
}

func validatePriceFeed(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("price feed cannot be blank")
	}

	return nil
}

func validateMaxContractAgeInBlocks(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return errors.New("max contract age in blocks cannot be negative")
	}

	return nil
}

func validatePricePerTbPerMonth(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return errors.New("price per tb per month cannot be negative")
	}

	return nil
}
