package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMintDenom      = []byte("MintDenom")
	KeyProviderRatio  = []byte("ProviderRatio")
	KeyTokensPerBlock = []byte("TokensPerBlock")

	// TODO: Determine the default value
	DefaultMintDenom      = "ujkl"
	DefaultProviderRatio  = int64(4)
	DefaultTokensPerBlock = int64(6_000_000)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	mintDenom string,
	providerRatio int64,
	tokensPerBlock int64,
) Params {
	return Params{
		MintDenom:      mintDenom,
		ProviderRatio:  providerRatio,
		TokensPerBlock: tokensPerBlock,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMintDenom,
		DefaultProviderRatio,
		DefaultTokensPerBlock,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyProviderRatio, &p.ProviderRatio, validateProviderRatio),
		paramtypes.NewParamSetPair(KeyTokensPerBlock, &p.TokensPerBlock, validateTokensPerBlock),
	}
}

// Validate validates the set of params
func (p *Params) Validate() error {
	err := validateMintDenom(p.MintDenom)
	if err != nil {
		return err
	}
	err = validateTokensPerBlock(p.TokensPerBlock)
	if err != nil {
		return err
	}

	err = validateProviderRatio(p.ProviderRatio)
	if err != nil {
		return err
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

// validateTokensPerBlock validates the TokensMintedPerBlock param
func validateTokensPerBlock(v interface{}) error {
	tokensPerBlock, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if tokensPerBlock < 0 {
		return fmt.Errorf("must be greater or equal to 0: %T", v)
	}

	return nil
}

// validateProviderRatio validates the ProviderRatio param
func validateProviderRatio(v interface{}) error {
	ratio, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if ratio < 0 {
		return fmt.Errorf("must be greater or equal to 0: %T", v)
	}

	if ratio > 10 {
		return fmt.Errorf("must be less than or equal to 10: %T", v)
	}

	return nil
}
