package types

import (
	"errors"
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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
	KeyAttestFormSize         = []byte("AttestFormSize")
	KeyAttestMinToPass        = []byte("AttestMinToPass")
	KeyCollateralPrice        = []byte("CollateralPrice")
	KeyCheckWindow            = []byte("CheckWindow")
	KeyReferrals              = []byte("Referrals")
	KeyPOLRatio               = []byte("POLRatio")

	DefaultReferrals = int64(25)
	DefaultPOLRatio  = int64(40)
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
		AttestMinToPass:        3,
		AttestFormSize:         5,
		CollateralPrice:        10_000_000_000,
		CheckWindow:            100,
		ReferralCommission:     DefaultReferrals,
		PolRatio:               DefaultPOLRatio,
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
		paramtypes.NewParamSetPair(
			KeyCheckWindow,
			&p.CheckWindow,
			validateCheckWindow),

		paramtypes.NewParamSetPair(KeyPOLRatio, &p.PolRatio, validateInt64),
		paramtypes.NewParamSetPair(KeyReferrals, &p.ReferralCommission, validateInt64),
	}
}

func validateCheckWindow(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 1 {
		return errors.New("check window must be greater than than 1")
	}

	return nil
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
func (p *Params) Validate() error {
	err := validateInt64(p.PolRatio)
	if err != nil {
		return sdkerrors.Wrapf(err, "pol ratio is %d", p.PolRatio)
	}

	err = validateInt64(p.ReferralCommission)
	if err != nil {
		return sdkerrors.Wrapf(err, "referral commission is %d", p.ReferralCommission)
	}

	return nil
}

// validateInt validates the param is an int64
func validateInt64(v interface{}) error {
	i, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if i < 0 {
		return fmt.Errorf("must be greater or equal to 0 but is %d: %T", i, v)
	}

	return nil
}

// String implements the Stringer interface.
func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
