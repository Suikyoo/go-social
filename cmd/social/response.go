package main

import (
	"net/http"
)

//these are the common response codes that I usually end up sending to the clients

//it's a system error no the backend, therefore for easier problem diagnosis, 
//just return the error message of what caused the error
func InternalError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

//most client errors are wrong payloads so I'm just gonna put it as BadRequest
//I'll add more functions if I encounter client errors other than this
func RequestError(w http.ResponseWriter, err_text string) {
	http.Error(w, err_text, http.StatusBadRequest)
}

//this would've been an internal error if it weren't for the fact that I can't just
//reveal database error logs to clients due to security risks
func DBError(w http.ResponseWriter) {
	http.Error(w, "Database Error", http.StatusInternalServerError)
}

func AuthError(w http.ResponseWriter) {
	http.Error(w, "User not authorized", http.StatusUnauthorized)
}
