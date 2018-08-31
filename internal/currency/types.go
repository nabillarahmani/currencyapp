package currency

import (
	"fmt"

	"github.com/nabillarahmani/currencyapp/internal/common/database"
)

var (
	// DBConn is a main database object for currency
	DBConn database.Database
)

type (
	//CurrIface is an interface to interact with currency operations
	CurrIface interface {
		// for currency
		InsertCurrency(CurrData) error
		GetCurrency() error
		IsCurrencyExist(CurrData) (bool, error)
		// for currency rates
		InsertCurrencyRates() error
		GetCurrencyRates(string, string) error
		GetCurrencyRatesByDate(string) error
		RemoveCurrencyRates(string, string, string) error
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

	// ArrCurrData is a container for many CurrData
	ArrCurrData struct {
		Datas []CurrData `json:"currency_datas"`
	}
)

var (
	// ErrCurrencyExist is error for describing currency already exist
	ErrCurrencyExist = fmt.Errorf("Error currency already exist")
)
