package main

import (
	"net/http"
	"strconv"
)

var Count = 0

func IncreaseCount(w http.ResponseWriter, r *http.Request) {
	Count++

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(Count)))
}

func DecrementCount(w http.ResponseWriter, r *http.Request) {
	Count--

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(Count)))
}
