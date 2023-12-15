package v5

import (
	"errors"
	"fmt"
	"strings"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*v4Params)(nil)

// Params defines the parameters for the module.
type v4Params struct {
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
	AttestFormSize         int64 `protobuf:"varint,8,opt,name=attestFormSize,proto3" json:"attestFormSize,omitempty"`
	AttestMinToPass        int64 `protobuf:"varint,9,opt,name=attestMinToPass,proto3" json:"attestMinToPass,omitempty"`
	CollateralPrice        int64 `protobuf:"varint,10,opt,name=collateralPrice,proto3" json:"collateralPrice,omitempty"`
}

var (
	KeyDepositAccount         = []byte("DepositAccount")
	KeyProofWindow            = []byte("ProofWindow")
	KeyChunkSize              = []byte("ChunkSize")
	KeyMissesToBurn           = []byte("MissesToBurn")
	KeyPriceFeed              = []byte("PriceFeed")
	KeyMaxContractAgeInBlocks = []byte("MaxContractAgeInBlocks")
	KeyPricePerTbPerMonth     = []byte("PricePerTbPerMonth")
	KeyAttestFormSize         = []byte("AttestFormSize")
	KeyAttestMinToPass        = []byte("AttestMinToPass")
	KeyCollateralPrice        = []byte("CollateralPrice")
)

// ParamSetPairs get the params.ParamSet
func (p *v4Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDepositAccount, &p.DepositAccount, validateDeposit),
		paramtypes.NewParamSetPair(KeyProofWindow, &p.ProofWindow, validateProofWindow),
		paramtypes.NewParamSetPair(KeyChunkSize, &p.ChunkSize, validateChunkSize),
		paramtypes.NewParamSetPair(KeyMissesToBurn, &p.MissesToBurn, validateMissesToBurn),
		paramtypes.NewParamSetPair(KeyPriceFeed, &p.PriceFeed, validatePriceFeed),
		paramtypes.NewParamSetPair(
			KeyMaxContractAgeInBlocks,
			&p.MaxContractAgeInBlocks,
			validateMaxContractAgeInBlocks),
		paramtypes.NewParamSetPair(
			KeyPricePerTbPerMonth,
			&p.PricePerTbPerMonth,
			validatePricePerTbPerMonth),
		paramtypes.NewParamSetPair(
			KeyAttestFormSize,
			&p.AttestFormSize,
			validateAttestFormSize),
		paramtypes.NewParamSetPair(
			KeyAttestMinToPass,
			&p.AttestMinToPass,
			validateAttestMinToPass),
		paramtypes.NewParamSetPair(
			KeyCollateralPrice,
			&p.CollateralPrice,
			validateCollateralPrice),
	}
}

func validateCollateralPrice(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 1 {
		return errors.New("collateral price must be greater than 1")
	}

	return nil
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

func validateAttestMinToPass(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return errors.New("min to pass cannot be negative")
	}

	return nil
}

func validateAttestFormSize(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return errors.New("form size cannot be negative")
	}

	return nil
}

// Validate validates the set of params
func (p *v4Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p *v4Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
