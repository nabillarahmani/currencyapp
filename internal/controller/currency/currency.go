package currency

import (
	"database/sql"
	"strconv"
	"time"

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

// isValidParam is to check whether param is valid or not
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
	result, err = pkg.IsCurrencyExist(currData.From, currData.To)
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

// AddCurrencyRatesController is a controller for add new currency rates data
func AddCurrencyRatesController(data RatesData) (result bool, err error) {
	// validate parameter, if valid then true
	result, err = isValidParamRates(data)
	if !result {
		log.Error(err)
		return
	}

	// init pkg
	pkg := currency.CurrPkg{}

	// prepare Data
	var currData currency.CurrRatesData
	currData.From = data.From
	currData.To = data.To
	currData.Rates, _ = strconv.ParseFloat(data.Rates, 64)
	currData.Date = data.Date

	// check if currency exist, if not exist then do not allow to insert rates
	result, err = pkg.IsCurrencyExist(currData.From, currData.To)
	if !result {
		err = ErrCurrencyNotExist
		log.Error(err)
		return
	}

	// if does exist, just upsert data to table rates
	err = pkg.UpsertCurrencyRates(currData)
	if err != nil {
		log.Error(err, "error while upsert currency rates")
		return
	}

	return
}

// isValidParamRates is to check whether param is valid or not
func isValidParamRates(data RatesData) (result bool, err error) {
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

	// check date related
	result, err = isValidDate(data.Date)
	if err != nil {
		return
	}

	// check rates
	_, err = strconv.ParseFloat(data.Rates, 64)
	if err != nil {
		log.Error(err)
		result = false
		err = ErrRatesInvalid
		return
	}

	result = true
	return
}

// isValidDate to check date
func isValidDate(date string) (bool, error) {
	if len(date) == 10 {
		// check if given date parameter is in right format
		timeDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			return false, ErrDateFormatInvalid
		}
		// check if given date is after now, if so return invalid
		if timeDate.After(time.Now()) {
			return false, ErrDateInvalid
		}
	} else {
		return false, ErrDateLengthInvalid
	}
	return true, nil
}

// GetCurrencyRatesController is a controller to get all currency rates for past 7 days
func GetCurrencyRatesController(date string) (data currency.ArrWrapperCurrRatesData, err error) {
	// validate date
	flag, err := isValidDate(date)
	if err != nil && flag == false {
		log.Errorf(err, "error when validate date, with date:[%s]", date)
		return
	}

	// get valid date for past 7 days
	sdate := getOneWeekPrevPeriode(date)

	// get all currency
	arrCurrData, err := GetAllCurrencyController()
	if err != nil {
		log.Error(err, "error when retrieve all currency")
		return
	}

	// init pkg
	pkg := currency.CurrPkg{}

	// loop for each currency data
	for _, obj := range arrCurrData.Data {
		var tempRes currency.WrapperCurrRatesData
		// get currency rates data with specified currency and date range
		tempRes, err = pkg.GetCurrencyRatesByDate(obj.From, obj.To, sdate, date)
		if err != nil {
			log.Error(err, "error when retrieve currency rates data by date")
			return
		}

		data.Data = append(data.Data, tempRes)
	}
	return
}

// getOneWeekPrevPeriode is a function to get one week prev
func getOneWeekPrevPeriode(date string) string {
	var sdate string

	dateTime, _ := time.Parse("2006-01-02", date)

	sdate = dateTime.AddDate(0, 0, -6).Format("2006-01-02")

	return sdate
}

// GetCurrencyRatesTrendController is a controller for getting all currency rates tren
func GetCurrencyRatesTrendController() (err error) {
	return
}
