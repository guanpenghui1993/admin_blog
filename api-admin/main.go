package main

import (
	"api-admin/helpers"
	"api-admin/router"
	"net/http"
	"time"

	"fmt"
)

func main() {

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", helpers.ServerConfig.Port),
		Handler:        router.Route,
		ReadTimeout:    helpers.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout:   helpers.ServerConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
