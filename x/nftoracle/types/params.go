package types

import (
	"fmt"
	time "time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var (
	KeyMintRequestActiveDuration = []byte("MintRequestActiveDuration")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(mintRequestActiveDuration time.Duration) Params {
	return Params{
		MintRequestActiveDuration: mintRequestActiveDuration,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	defaultDuration := 0 * time.Microsecond
	return NewParams(defaultDuration)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintRequestActiveDuration, &p.MintRequestActiveDuration, validateMintRequestActiveDuration)}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMintRequestActiveDuration(p.MintRequestActiveDuration); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateMintRequestActiveDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}
