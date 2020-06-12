package common

import (
	"encoding/json"
	"net/http"
)

// ResJSON response with JSON with payload stream
func ResJSON(w http.ResponseWriter, httpCode int, payload interface{}) (err error) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)
	err = json.NewEncoder(w).Encode(payload)
	return
}

// ResJSONWithData response with data,
// HTTP status will be 200, and message = "successfull retrieve some data"
// If you want to custom http code and message, use ResJSON instead
func ResJSONWithData(w http.ResponseWriter, data interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var resp ResponseWithData
	resp.Message = "successful retrieve some data"
	resp.Data = data
	err = json.NewEncoder(w).Encode(resp)
	return
}

//ResJSONErrorField response error when error location in field
func ResJSONErrorField(w http.ResponseWriter, data interface{}) (err error) {
	w.WriteHeader(422)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var resp ResponseErrField
	resp.Message = "Some field data not suitable the requirement"
	resp.ErrorType = "error_fields"
	resp.Field = data
	err = json.NewEncoder(w).Encode(resp)
	return
}

//ResponseNoData type response with no data
type ResponseNoData struct {
	Message string `json:"message"`
}

//ResponseWithData type response with data
type ResponseWithData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//ResponseErrField error field type response with data
type ResponseErrField struct {
	Message   string      `json:"message"`
	ErrorType string      `json:"error_type`
	Field     interface{} `json:"field"`
}
