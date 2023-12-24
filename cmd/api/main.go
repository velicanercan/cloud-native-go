package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/velicanercan/cloud-native-go/cmd/api/router"
	"github.com/velicanercan/cloud-native-go/util/validator"
	"github.com/velicanercan/cloud-native-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const fmtDbString = "host=%s port=%d user=%s dbname=%s password=%s sslmode=disable"

//	@title			Cloud Native Go API
//	@version		1.0
//	@description	This is a sample server for Cloud Native Go Project.

//	@contact.name	Velican Ercan
//	@contact.url	github/velicanercan

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	c := config.New()
	v := validator.New()
	var logLevel gormlogger.LogLevel
	switch c.DB.Debug {
	case true:
		logLevel = gormlogger.Info
	default:
		logLevel = gormlogger.Error
	}

	dbString := fmt.Sprintf(fmtDbString, c.DB.Host, c.DB.Port, c.DB.Username, c.DB.DBName, c.DB.Password)

	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	if err != nil {
		slog.Error("Failed to connect database", "err", err)
		return
	}

	r := router.New(db, v)

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
