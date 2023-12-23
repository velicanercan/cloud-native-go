package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/velicanercan/cloud-native-go/config"
)

func main() {
	c := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
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
