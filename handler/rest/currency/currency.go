package currency

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nabillarahmani/currencyapp/internal/common/log"
	"github.com/nabillarahmani/currencyapp/internal/common/response"
	"github.com/nabillarahmani/currencyapp/internal/controller/currency"
)

// AddRemoveCurrency is a handler for add/remove currency
func AddRemoveCurrency(w http.ResponseWriter, r *http.Request) {
	// prepare response result
	var resp response.HandlerResponse
	// get data from body
	var data currency.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Error(err, "There's an error while decode the body")
		resp.StatusCode = http.StatusBadRequest
		response.WriteAPIStandard(w, resp, err)
		return
	}

	// validation data
	if data.From == "" {
		err = fmt.Errorf("From parameter is empty")
		log.Error(err)
		resp.StatusCode = http.StatusBadRequest
		response.WriteAPIStandard(w, resp, err)
		return
	}
	if data.To == "" {
		err = fmt.Errorf("To parameter is empty")
		log.Error(err)
		resp.StatusCode = http.StatusBadRequest
		response.WriteAPIStandard(w, resp, err)
		return
	}
	if data.Action == "" {
		err = fmt.Errorf("Action parameter is empty")
		log.Error(err)
		resp.StatusCode = http.StatusBadRequest
		response.WriteAPIStandard(w, resp, err)
		return
	}

	flagCont := currency.IsValidAction(data.Action)
	if !flagCont {
		err = currency.ErrInvalidParamAction
		log.Error(err)
		resp.StatusCode = http.StatusBadRequest
		response.WriteAPIStandard(w, resp, err)
		return
	}

	// proceed data
	var res bool
	switch data.Action {
	case "add":
		res, err = currency.AddCurrencyController(data)
	case "remove":
		res, err = currency.RemoveCurrencyController(data)
	}
	if err != nil {
		if err == currency.ErrCurrencyExist {
			resp.StatusCode = http.StatusBadRequest
		} else if err == currency.ErrCurrencyNotExist {
			resp.StatusCode = http.StatusBadRequest
		} else if err == currency.ErrInvalidParamFrom {
			resp.StatusCode = http.StatusBadRequest
		} else if err == currency.ErrInvalidParamTo {
			resp.StatusCode = http.StatusBadRequest
		} else {
			resp.StatusCode = http.StatusInternalServerError
		}
		resp.Message = res
		log.Error(err)
		response.WriteAPIStandard(w, resp, err)
		return
	}

	// write to response
	resp.Message = res
	resp.StatusCode = http.StatusOK
	response.WriteAPIStandard(w, resp, nil)
	return
}

// GetCurrency is a handler for get all currencies
func GetCurrency(w http.ResponseWriter, r *http.Request) {
	// prepare response result
	var resp response.HandlerResponse

	// proceed data
	res, err := currency.GetAllCurrencyController()
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		log.Error(err)
		response.WriteAPIStandard(w, resp, err)
		return
	}

	// write to response
	resp.Message = res
	resp.StatusCode = http.StatusOK
	response.WriteAPIStandard(w, resp, nil)
	return
}
