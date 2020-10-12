package main

import (
	"fmt"
	"net/http"
	"time"
	"watt/pkg/router"
	"watt/pkg/utils"
)

func main() {

	// 这块协程有个bug
	// if utils.Setting.Common.AutoTable {
	// 	go models.AutoMigrateTable()
	// }

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", utils.Setting.Http.Port),
		Handler:        router.Route,
		ReadTimeout:    utils.Setting.Http.ReadTimeout * time.Second,
		WriteTimeout:   utils.Setting.Http.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()

	if err != nil {
		utils.Exit(err)
	}
}
