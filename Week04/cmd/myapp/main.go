package main

import (
	"fmt"
	setting "go-note/Week04/internal/pkg"
	"go-note/Week04/internal/service"
	"net/http"
)

func main() {
	router := service.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
