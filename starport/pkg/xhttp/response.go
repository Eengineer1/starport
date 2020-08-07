package xhttp

import (
	"encoding/json"
	"errors"
	"net/http"
)

// ResponseJSON writes a JSON response to w by using status as http status and data
// as payload.
func ResponseJSON(w http.ResponseWriter, status int, data interface{}) error {
	bdata, err := json.Marshal(data)
	if err != nil {
		status = http.StatusInternalServerError
		bdata, _ = json.Marshal(NewErrorResponse(errors.New(http.StatusText(status))))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bdata)
	return err
}

// ErrorResponseBody is the skeleton for error messages that should be sent to
// client.
type ErrorResponseBody struct {
	Error ErrorResponse `json:"error"`
}

// ErrorResponse holds the error message.
type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse creates a new http error response from err.
func NewErrorResponse(err error) ErrorResponseBody {
	return ErrorResponseBody{
		Error: ErrorResponse{
			Message: err.Error(),
		},
	}
}
