package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type configuration struct {
	port        int
	environment string
}

type application struct {
	config configuration
	logger *slog.Logger
}

func main() {
	var config configuration

	flag.IntVar(&config.port, "port", 4000, "Api server port")
	flag.StringVar(&config.environment, "env", "development", "The Environment the app is running in")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: config,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", config.environment)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
