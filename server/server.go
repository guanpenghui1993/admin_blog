package server

import (
	"fmt"
	"net/http"
	"time"
	"watt/pkg/config"
	"watt/pkg/helpers"
	"watt/pkg/router"
)

func Run() {

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Conf.Http.Port),
		Handler:        router.Route,
		ReadTimeout:    config.Conf.Http.ReadTimeout * time.Second,
		WriteTimeout:   config.Conf.Http.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()

	if err != nil {
		helpers.H.Exit(err)
	}
}
