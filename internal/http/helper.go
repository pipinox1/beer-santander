package http

import (
	"beer/internal/tools/customerror"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorBody struct {
	Message string
	Status  int
	Code    string
}

// respondwithError return error message
func ErrorResponse(w http.ResponseWriter, err error) {
	errorHandled := ErrorHandler(err)
	WebResponse(w, errorHandled.Status,errorHandled )
}

// WebResponse write json response format
func WebResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func ErrorHandler(err error) ErrorBody {
	switch err.(type) {
	case customerror.BusinessError:
		return ErrorBody{err.Error(), http.StatusBadRequest, "bad_request"}
	case customerror.ExternalServiceError:
		return ErrorBody{err.Error(), http.StatusServiceUnavailable, "service_unavailable"}
	}
	return ErrorBody{"internal_server_error", http.StatusInternalServerError, "internal_error"}
}
