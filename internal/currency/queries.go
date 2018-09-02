package currency

const (
	queryUpsertCurrency = `
		INSERT INTO ws_currency (
			from_curr,
			to_curr,
			status
		) VALUES (
			$1,
			$2,
			$3
		) ON CONFLICT(from_curr, to_curr)
			DO UPDATE SET status = $3
	`

	queryCheckCurrency = `
		SELECT 
			status
		FROM
			ws_currency
		WHERE
			from_curr = $1 AND
			to_curr = $2
		LIMIT 1
	`

	queryGetAllCurrency = `
		SELECT 
			from_curr, to_curr, status
		FROM
			ws_currency
		WHERE status = 1
	`

	queryUpsertRatesCurrency = `
		INSERT INTO ws_currency_rates (
			date,
			from_curr,
			to_curr,
			rates
		) VALUES (
			$1,
			$2,
			$3,
			$4
		) ON CONFLICT(date, from_curr, to_curr)
			DO UPDATE SET rates = $4
	`

	queryGetCurrencyRates = `
		SELECT
			date, rates
		FROM 
			ws_currency_rates
		WHERE
			date <= $1 
				AND 
			date >= $2
				AND
			from_curr = $3
				AND
			to_curr=  $4
	`

	queryGetLatest7CurrencyRates = `
		SELECT
			date, rates
		FROM
			ws_currency_rates
		WHERE
			from_curr = $1
				AND
			to_curr = $2
		LIMIT 7
	`
)
