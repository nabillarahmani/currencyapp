package currency

import (
	"fmt"
)

const (
	// Inactive means deleted
	Inactive = 0
	// Active means exist
	Active = 1
)

type (
	// Data is a struct for currency data
	Data struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Action string `json:"action"`
	}
)

var (
	// ErrCurrencyExist is error for describing currency already exist
	ErrCurrencyExist = fmt.Errorf("Error currency already exist")
	// ErrCurrencyNotExist is error for describing currency not exist
	ErrCurrencyNotExist = fmt.Errorf("Error currency not exist")
	// ErrInvalidParamFrom is error for invalid param
	ErrInvalidParamFrom = fmt.Errorf("Error invalid param from")
	// ErrInvalidParamTo is error for invalid param
	ErrInvalidParamTo = fmt.Errorf("Error invalid param to")
	// ErrInvalidParamAction is error for invalid param
	ErrInvalidParamAction = fmt.Errorf("Error invalid param action")
)
