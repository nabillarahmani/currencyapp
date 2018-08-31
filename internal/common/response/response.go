package response

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type (
	//HandlerResponse is standarized response type for all api handler
	HandlerResponse struct {
		Message    interface{} // data
		StatusCode int         // status code
	}

	// RespAPIStandard is response standard
	RespAPIStandard struct {
		Header struct {
			Reason     string `json:"reason"`
			StatusCode string `json:"status_code"`
		} `json:"header"`
		Data interface{} `json:"data"`
	}
)

// WriteAPIStandard writes header based on RespAPIStandard
func WriteAPIStandard(w http.ResponseWriter, dataResp HandlerResponse, err error) {
	var response RespAPIStandard
	if err != nil {
		response.Header.Reason = err.Error()
	}
	response.Header.StatusCode = strconv.Itoa(dataResp.StatusCode)
	response.Data = dataResp.Message

	// marshall and write response
	e, _ := json.Marshal(response)
	if w.Header().Get("Access-Control-Allow-Origin") == "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(dataResp.StatusCode)
	w.Write(e)
}
