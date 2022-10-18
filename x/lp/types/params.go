package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

/*
Params structure:
type Params struct {
	// Contains proto Params message parameters at param.proto.
}
*/

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter keys
const (
	KeyMinInitPoolDeposit = "MinInitPoolDeposit"
	KeyMaxPoolDenomCount  = "MaxPoolDenomCount"
	KeyLPTokenUnit        = "LPTokenUnit"
)

// Default values
var (
	DefaultInitPoolDeposit   uint64 = 2
	DefaultMaxPoolDenomCount uint32 = 2
	DefaultLPTokenUnit       uint32 = 6
)

// Validation methods
func validateMinInitPoolDeposit(i interface{}) error {
	// Type assertion.
	_, ok := i.(uint64)

	if !ok {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidType,
			"Parameter validation error at %s module, %s must be uint64",
			ModuleName,
			KeyMinInitPoolDeposit,
		)
	}

	return nil
}

func validateMaxPoolDenomCount(i interface{}) error {
	// Type assertion.
	_, ok := i.(uint32)

	if !ok {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidType,
			"Parameter validation error at %s module, %s must be uint32",
			ModuleName,
			KeyMaxPoolDenomCount,
		)
	}

	value, _ := i.(uint32)

	if value < 1 {
		return sdkerrors.Wrapf(
			ErrInvalidValue,
			"Parameter (%s) validation error at %s module, denom count must be bigger than 1",
			KeyMaxPoolDenomCount,
			ModuleName,
		)
	}

	return nil
}

func validateLPTokenUnit(i interface{}) error {
	// Type assertion.
	_, ok := i.(uint32)

	if !ok {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidType,
			"Parameter validation error at %s module, %s must be uint32",
			ModuleName,
			KeyMaxPoolDenomCount,
		)
	}

	return nil
}

// Initialize param keytable at module launch.
// Default param values are used for initialization.
// This is used at NewKeeper() in this module.
func ParamKeyTable() paramtypes.KeyTable {
	defaultParams := DefaultParams()
	return paramtypes.NewKeyTable().RegisterParamSet(&defaultParams)
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		MinInitPoolDeposit: DefaultInitPoolDeposit,
		MaxPoolDenomCount:  DefaultMaxPoolDenomCount,
		LPTokenUnit:        DefaultLPTokenUnit,
	}
}

/*
ParamSetPair structure:
	type ParamSetPair struct {
		Key         []byte
		Value       interface{}
		ValidatorFn ValueValidatorFn
	}

source:
	https://pkg.go.dev/github.com/cosmos/cosmos-sdk@v0.45.6/x/params/types#ParamSetPair
*/

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	// return slice of ParamSetPair
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair([]byte(KeyMinInitPoolDeposit), &p.MinInitPoolDeposit, validateMinInitPoolDeposit),
		paramtypes.NewParamSetPair([]byte(KeyMaxPoolDenomCount), &p.MaxPoolDenomCount, validateMaxPoolDenomCount),
		paramtypes.NewParamSetPair([]byte(KeyLPTokenUnit), &p.LPTokenUnit, validateLPTokenUnit),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
