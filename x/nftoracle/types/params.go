package types

import (
	"fmt"
	time "time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	// DefaultParamspace default name for parameter store
	DefaultMintRequestActiveDuration   = 120 * time.Second
	DefaultActionRequestActiveDuration = 120 * time.Second
	DefaultVerifyRequestActiveDuration = 120 * time.Second
)

var (
	KeyMintRequestActiveDuration   = []byte("MintRequestActiveDuration")
	KeyActionRequestActiveDuration = []byte("ActionRequestActiveDuration")
	KeyVerifyRequestActiveDuration = []byte("VerifyRequestActiveDuration")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(mintRequestActiveDuration time.Duration, actionRequestActiveDuration time.Duration, verifyRequestActiveDuration time.Duration) Params {
	return Params{
		MintRequestActiveDuration:   mintRequestActiveDuration,
		ActionRequestActiveDuration: actionRequestActiveDuration,
		VerifyRequestActiveDuration: verifyRequestActiveDuration,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMintRequestActiveDuration, DefaultActionRequestActiveDuration, DefaultVerifyRequestActiveDuration)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintRequestActiveDuration, &p.MintRequestActiveDuration, validateMintRequestActiveDuration),
		paramtypes.NewParamSetPair(KeyActionRequestActiveDuration, &p.ActionRequestActiveDuration, validateActionRequestActiveDuration),
		paramtypes.NewParamSetPair(KeyVerifyRequestActiveDuration, &p.VerifyRequestActiveDuration, validateVerifyRequestActiveDuration),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMintRequestActiveDuration(p.MintRequestActiveDuration); err != nil {
		return err
	}
	if err := validateActionRequestActiveDuration(p.ActionRequestActiveDuration); err != nil {
		return err
	}
	if err := validateVerifyRequestActiveDuration(p.VerifyRequestActiveDuration); err != nil {
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

func validateActionRequestActiveDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}

func validateVerifyRequestActiveDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}
