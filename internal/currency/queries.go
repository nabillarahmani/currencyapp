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
	`
)
