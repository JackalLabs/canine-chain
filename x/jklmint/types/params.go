package types

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMintDenom      = []byte("MintDenom")
	KeyTokensPerBlock = []byte("TokensPerBlock")
	KeyDevGrants      = []byte("DevGrants")
	KeyStakerRatio    = []byte("StakerRatio")
	KeyMintIncrease   = []byte("MintIncrease")
	KeyStorageStipend = []byte("StorageStipend")
	KeyProviderRatio  = []byte("ProviderRatio")

	// TODO: Determine the default value
	DefaultMintDenom      = "ujkl"
	DefaultTokensPerBlock = int64(4_200_000)
	DefaultDevGrants      = int64(8)
	DefaultStakerRatio    = int64(80)
	DefaultMintDecrease   = int64(6)
	DefaultStorageStipend = "jkl18dtaqkj3cdazn4rpgqdc3acz98cp5yz30erp95"
	DefaultProviderRatio  = int64(12)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	mintDenom string,
	devGrants int64,
	tokensPerBlock int64,
	stakerRatio int64,
	mintDecrease int64,
	storageStipendAddress string,
	storageProviderRatio int64,
) Params {
	return Params{
		MintDenom:             mintDenom,
		DevGrantsRatio:        devGrants,
		StakerRatio:           stakerRatio,
		TokensPerBlock:        tokensPerBlock,
		MintDecrease:          mintDecrease,
		StorageStipendAddress: storageStipendAddress,
		StorageProviderRatio:  storageProviderRatio,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMintDenom,
		DefaultDevGrants,
		DefaultTokensPerBlock,
		DefaultStakerRatio,
		DefaultMintDecrease,
		DefaultStorageStipend,
		DefaultProviderRatio,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyTokensPerBlock, &p.TokensPerBlock, validateInt64),
		paramtypes.NewParamSetPair(KeyDevGrants, &p.DevGrantsRatio, validateInt64),
		paramtypes.NewParamSetPair(KeyMintIncrease, &p.MintDecrease, validateInt64),
		paramtypes.NewParamSetPair(KeyStakerRatio, &p.StakerRatio, validateInt64),
		paramtypes.NewParamSetPair(KeyStorageStipend, &p.StorageStipendAddress, validateStipend),
		paramtypes.NewParamSetPair(KeyProviderRatio, &p.StorageProviderRatio, validateInt64),
	}
}

// Validate validates the set of params
func (p *Params) Validate() error {
	err := validateInt64(p.StorageProviderRatio)
	if err != nil {
		return sdkerrors.Wrapf(err, "storage p ratio is %d", p.StorageProviderRatio)
	}

	err = validateMintDenom(p.MintDenom)
	if err != nil {
		return err
	}

	err = validateStipend(p.StorageStipendAddress)
	if err != nil {
		return err
	}

	err = validateInt64(p.TokensPerBlock)
	if err != nil {
		return sdkerrors.Wrapf(err, "tokens per block is %d", p.TokensPerBlock)
	}

	err = validateInt64(p.DevGrantsRatio)
	if err != nil {
		return sdkerrors.Wrapf(err, "dev grants ratio is %d", p.DevGrantsRatio)
	}

	err = validateInt64(p.MintDecrease)
	if err != nil {
		return sdkerrors.Wrapf(err, "mint decrease is %d", p.MintDecrease)
	}

	err = validateInt64(p.StakerRatio)
	if err != nil {
		return sdkerrors.Wrapf(err, "staker ratio is %d", p.StakerRatio)
	}
	return nil
}

// String implements the Stringer interface.
func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMintDenom validates the MintDenom param
func validateMintDenom(v interface{}) error {
	mintDenom, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = mintDenom

	return nil
}

// validateMintDenom validates the MintDenom param
func validateStipend(v interface{}) error {
	_, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
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
