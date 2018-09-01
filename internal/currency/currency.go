package currency

import (
	"strconv"
	"time"

	"github.com/nabillarahmani/currencyapp/internal/common/database"
	"github.com/nabillarahmani/currencyapp/internal/common/log"
)

// InitPackage will store the db object so we wont need to call it from global var config
func InitPackage(dbobj database.Database) {
	DBConn = dbobj
}

// UpsertCurrency is a method for insert/update new currency
func (pkg CurrPkg) UpsertCurrency(data CurrData) (err error) {
	// prepare tx
	tx, err := DBConn.Begin()
	if err != nil {
		log.Error(err)
		return
	}

	// defer
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
			log.Errorf(err, "Error when exec with data:%+v", data)
		}
	}()

	// exec upsert
	_, err = tx.Exec(
		queryUpsertCurrency,
		data.From,
		data.To,
		data.Status,
	)

	return
}

// IsCurrencyExist will return flag if currency exist
func (pkg CurrPkg) IsCurrencyExist(from, to string) (flag bool, err error) {
	// get result
	rows, err := DBConn.Queryx(queryCheckCurrency, from, to)
	if err != nil {
		log.Errorf(err, "error when querying all currency data, with from:[%s], to:[%s]", from, to)
		return
	}
	defer rows.Close()

	flag = rows.Next()
	return
}

// GetCurrency will return all currency
func (pkg CurrPkg) GetCurrency() (res ArrCurrData, err error) {
	// get result
	rows, err := DBConn.Queryx(queryGetAllCurrency)
	if err != nil {
		log.Errorf(err, "error when querying all currency data, with data:[%+v]", res)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var temp CurrData
		err = rows.Scan(
			&temp.From,
			&temp.To,
			&temp.Status,
		)
		if err != nil {
			log.Error(err, "Error when scan data")
			return
		}

		res.Data = append(res.Data, temp)
	}

	return
}

// UpsertCurrencyRates will upsert data
func (pkg CurrPkg) UpsertCurrencyRates(data CurrRatesData) (err error) {
	// prepare tx
	tx, err := DBConn.Begin()
	if err != nil {
		log.Error(err)
		return
	}

	// defer
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
			log.Errorf(err, "Error when exec with data:%+v", data)
		}
	}()

	// exec upsert
	_, err = tx.Exec(
		queryUpsertRatesCurrency,
		data.Date,
		data.From,
		data.To,
		data.Rates,
	)

	return
}

// GetCurrencyRatesByDate will return currency rates for specific currency
func (pkg CurrPkg) GetCurrencyRatesByDate(from, to, sDate, eDate string) (res WrapperCurrRatesData, err error) {
	// get result
	rows, err := DBConn.Queryx(queryGetCurrencyRates, eDate, sDate, from, to)
	if err != nil {
		log.Errorf(err, "error when querying currency rate data, with param: from:[%s], to:[%s], sdate:[%s], edate:[%s]", from, to, sDate, eDate)
		return
	}
	defer rows.Close()

	res.From = from
	res.To = to
	var tempRates float64

	for rows.Next() {
		var tempDate time.Time
		var tempRate float64
		err = rows.Scan(
			&tempDate,
			&tempRate,
		)
		if err != nil {
			log.Error(err, "Error when scan data")
			return
		}
		if tempDate.Format("2006-01-02") == eDate {
			res.Rate = strconv.FormatFloat(tempRate, 'f', 5, 64)
		}
		tempRates += tempRate
	}

	// get 7 days average
	res.Avg = strconv.FormatFloat((tempRates / 7), 'f', 5, 64)
	return
}
