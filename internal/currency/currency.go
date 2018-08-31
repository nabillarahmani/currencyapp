package currency

import (
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
			log.Error(err)
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
func (pkg CurrPkg) IsCurrencyExist(res CurrData) (flag bool, err error) {
	// get result
	rows, err := DBConn.Queryx(queryCheckCurrency, res.From, res.To)
	if err != nil {
		log.Error(err, "error when querying all currency data")
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
		log.Error(err, "error when querying all currency data")

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
			log.Error(err, "Error when querying get all data")
			return
		}

		res.Datas = append(res.Datas, temp)
	}

	return
}

// InsertCurrencyRates will
func (pkg CurrPkg) InsertCurrencyRates() (err error) {
	return
}

// GetCurrencyRates will
func (pkg CurrPkg) GetCurrencyRates(from, to string) (err error) {
	return
}

// GetCurrencyRatesByDate will
func (pkg CurrPkg) GetCurrencyRatesByDate(date string) (err error) {
	return
}

// RemoveCurrencyRates will
func (pkg CurrPkg) RemoveCurrencyRates(date, from, to string) (err error) {
	return
}
