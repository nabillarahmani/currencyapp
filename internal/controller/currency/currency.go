package currency

import (
	"database/sql"

	"github.com/nabillarahmani/currencyapp/internal/common/log"
	"github.com/nabillarahmani/currencyapp/internal/currency"
)

// AddCurrencyController is a controller for adding new currency
func AddCurrencyController(data Data) (result bool, err error) {
	// validate parameter, if valid then true
	result, err = isValidParam(data)
	if !result {
		log.Error(err)
		return
	}

	// init pkg
	pkg := currency.CurrPkg{}

	// prepare Data
	var currData currency.CurrData
	currData.From = data.From
	currData.To = data.To
	currData.Status = Active

	// upsert data
	err = pkg.UpsertCurrency(currData)
	if err != nil {
		// update result to false for fail update
		result = false
		log.Error(err)
		return
	}

	return
}

// isValidParam is to check whether param is valid or nots
func isValidParam(data Data) (result bool, err error) {
	// check length of param From
	if len(data.From) > 3 || len(data.From) < 3 {
		result = false
		err = ErrInvalidParamFrom
		return
	}

	// check length of param To
	if len(data.To) > 3 || len(data.To) < 3 {
		result = false
		err = ErrInvalidParamTo
		return
	}
	result = true
	return
}

// IsValidAction will return flag for valid action
func IsValidAction(action string) bool {
	// check if action exist
	switch action {
	case "remove":
		return true
	case "add":
		return true
	default:
		return false
	}
}

// GetAllCurrencyController is a controller to get all data
func GetAllCurrencyController() (data currency.ArrCurrData, err error) {
	// init pkg
	pkg := currency.CurrPkg{}

	// get all datas
	data, err = pkg.GetCurrency()
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err, "error while retrieving data")
			return
		}
	}

	return
}

// RemoveCurrencyController is a controller for remove a currency
func RemoveCurrencyController(data Data) (result bool, err error) {
	// validate parameter, if valid then true
	result, err = isValidParam(data)
	if !result {
		log.Error(err)
		return
	}

	// init pkg
	pkg := currency.CurrPkg{}

	// prepare Data
	var currData currency.CurrData
	currData.From = data.From
	currData.To = data.To
	currData.Status = Inactive

	// check if currency exist
	result, err = pkg.IsCurrencyExist(currData)
	if !result {
		err = ErrCurrencyNotExist
		log.Error(err)
		return
	}

	// if does exist, just upsert data to inactive
	err = pkg.UpsertCurrency(currData)
	if err != nil {
		return
	}

	return
}
