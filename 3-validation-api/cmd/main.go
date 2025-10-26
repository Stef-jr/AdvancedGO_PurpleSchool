package main

import (
	"net/http"
	"validation/configs"
	"validation/internal"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	verify.NewVerifyHandler(router, verify.DepsVerify{
		Config: conf,
	})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	server.ListenAndServe()

}