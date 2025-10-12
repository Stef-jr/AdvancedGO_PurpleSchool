package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

type RandHandler struct{}

func NewRandHandler(router *http.ServeMux) {
	handler := &RandHandler{}
	router.HandleFunc("/rand", handler.RandHandler())
}

func (handler *RandHandler) RandHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := strconv.Itoa(rand.Intn(6))
		w.Write([]byte(n))
	}
}
