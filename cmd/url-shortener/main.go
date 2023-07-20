package main

import (
	"golang.org/x/exp/slog"
	"os"
	"petProject/internal/config"
	"petProject/internal/lib/logger/sl"
	"petProject/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	// log = log.With("env", cfg.Env) // adds env parameter to all logs

	log.Info("starting url-shortener", "env", cfg.Env)
	log.Debug("debug mode is on")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to create storage", sl.Err(err))
		os.Exit(1)
	}

	//_ = storage

	id, err := storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}

	log.Info("url saved", slog.Int64("id", id))

	id, err = storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}

	log.Info("url saved", slog.Int64("id", id))

	// TODO: inti router: chi, "chi reder"

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
