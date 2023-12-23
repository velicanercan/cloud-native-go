package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/velicanercan/cloud-native-go/cmd/api/router"
	"github.com/velicanercan/cloud-native-go/config"
)

//	@title			Cloud Native Go API
//	@version		1.0
//	@description	This is a sample server for Cloud Native Go Project.

//	@contact.name	Velican Ercan
//	@contact.url	github/velicanercan

//	@host		localhost:8080
//	@BasePath	/api/v1
func main() {
	c := config.New()
	r := router.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}
	slog.Info("Server started at", "port", strconv.Itoa(c.Server.Port))
	if err := s.ListenAndServe(); err != nil {
		slog.Error("Server stopped", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
