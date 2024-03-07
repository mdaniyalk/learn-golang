package api 

import (
	"encoding/json"
	"net/http"
)

// Coint Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	Code int

	Balance int64
}

type Error struct {
	Code int

	Message string
}


func writeError(w http.ResponseWriter, message string, code int) {
	response := Error{
		Code: code, 
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, "An Unexpected Error Occurred" + err.Error(), http.StatusInternalServerError)
	}
)