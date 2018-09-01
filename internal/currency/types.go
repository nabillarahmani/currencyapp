package currency

import (
	"fmt"
	"html/template"

	"github.com/nabillarahmani/currencyapp/internal/common/database"
)

var (
	// DBConn is a main database object for currency
	DBConn database.Database
	// addCurrencyTemplat is template for use case #4
	addCurrencyTemplate *template.Template
	// removeCurrencyTemplate is template for use case #5
	removeCurrencyTemplate *template.Template
	// addCurrencyRateTemplate is template for use case #1
	addCurrencyRateTemplate *template.Template
	// getCurrencyRateTemplate is template for use case #2
	getCurrencyRateTemplate *template.Template
	// trendCurrencyRateTemplate is template for use case #3
	trendCurrencyRateTemplate *template.Template
)

type (
	//CurrIface is an interface to interact with currency operations
	CurrIface interface {
		// for currency
		InsertCurrency(CurrData) error
		GetCurrency() error
		IsCurrencyExist(CurrData) (bool, error)
		// for currency rates
		UpsertCurrencyRates() error
		GetCurrencyRatesByDate(string) error
	}

	// CurrPkg is an object which implement CurrIface
	CurrPkg struct {
	}

	// CurrData is a struct for currency data
	CurrData struct {
		From   string `json:"from" db:"from_curr"`
		To     string `json:"to" db:"to_curr"`
		Status int    `json:"status" db:"status"`
	}

	// CurrRatesData is a struct for currency data
	CurrRatesData struct {
		Date  string  `json:"date" db:"date"`
		From  string  `json:"from" db:"from_curr"`
		To    string  `json:"to" db:"to_curr"`
		Rates float64 `json:"rates" db:"rates"`
	}

	// WrapperCurrRatesData is a wrapper for getting all currency rates data for past 7 days
	WrapperCurrRatesData struct {
		From string `json:"from"`
		To   string `json:"to"`
		Rate string `json:"rate"`
		Avg  string `json:"avg"` // 7 days avg
	}

	// ArrCurrData is a container for many CurrData
	ArrCurrData struct {
		Data []CurrData `json:"currency_data"`
	}

	// ArrWrapperCurrRatesData is a json response for currency rates data
	ArrWrapperCurrRatesData struct {
		Data []WrapperCurrRatesData `json:"wrapped_currency_rates_data"`
	}

	// ArrCurrRatesData is a container for many CurrRatesData
	ArrCurrRatesData struct {
		Data     []CurrRatesData `json:"currency_rates_data"`
		From     string          `json:"from"`
		To       string          `json:"to"`
		Avg      string          `json:"avg"`
		Variance string          `json:"variance"`
	}
)

var (
	// ErrCurrencyExist is error for describing currency already exist
	ErrCurrencyExist = fmt.Errorf("Error currency already exist")
)
