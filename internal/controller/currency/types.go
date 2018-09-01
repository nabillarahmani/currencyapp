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

	// RatesData is a struct for currency rates data
	RatesData struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Date  string `json:"date"`
		Rates string `json:"rates"`
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
	// ErrDateFormatInvalid is error for invalid date format
	ErrDateFormatInvalid = fmt.Errorf("Date format is invalid")
	// ErrDateLengthInvalid is error for length date
	ErrDateLengthInvalid = fmt.Errorf("Date length is invalid")
	// ErrDateInvalid is error for invalid date
	ErrDateInvalid = fmt.Errorf("Date is invalid to specify a request")
	// ErrRatesInvalid is error for invalid rates
	ErrRatesInvalid = fmt.Errorf("Rates invalid")
)
