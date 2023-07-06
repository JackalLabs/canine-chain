package types

import (
	"errors"
	"fmt"
	"strings"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyDepositAccount         = []byte("DepositAccount")
	KeyProofWindow            = []byte("ProofWindow")
	KeyChunkSize              = []byte("ChunkSize")
	KeyMissesToBurn           = []byte("MissesToBurn")
	KeyPriceFeed              = []byte("PriceFeed")
	KeyMaxContractAgeInBlocks = []byte("MaxContractAgeInBlocks")
	KeyPricePerTbPerMonth     = []byte("PricePerTbPerMonth")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		DepositAccount:         "cosmos1arsaayyj5tash86mwqudmcs2fd5jt5zgp07gl8",
		ProofWindow:            50,
		ChunkSize:              1024,
		MissesToBurn:           3,
		PriceFeed:              "jklprice",
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     8,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
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

// Validate validates the set of params
func (p *Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
