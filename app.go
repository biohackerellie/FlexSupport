package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	cfg "flexsupport/internal/config"
	"flexsupport/internal/lib/logger"
	"flexsupport/internal/router"
)

var log *slog.Logger

func App(ctx context.Context, stdout io.Writer, getenv func(string, string) string) error {
	config := cfg.New(getenv)
	local := config.Environment != cfg.PROD
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logLevel := config.LogLevel
	if logLevel == "" {
		logLevel = "debug"
	}
	logOptions := logger.LogOptions(strings.TrimSpace(strings.ToLower(logLevel)), config.VerboseLogging, local)
	switch local {
	case true:
		log = slog.New(slog.NewTextHandler(stdout, logOptions))
	default:
		log = slog.New(slog.NewJSONHandler(stdout, logOptions))
	}
	r := router.NewRouter(log)

	fmt.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", r)
}
