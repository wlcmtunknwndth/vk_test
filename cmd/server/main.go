package main

import (
	"context"
	"github.com/wlcmtunknwndth/vk_test/internal/app"
	"github.com/wlcmtunknwndth/vk_test/internal/config"
	"github.com/wlcmtunknwndth/vk_test/internal/lib/sl"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"

	scope = "main"
)

func main() {

	cfg := config.MustLoad(os.Getenv("config_path"))

	log := setupLogger(cfg.Env)

	log.Info("config", slog.Any("cfg", cfg))

	srv, err := app.New(context.Background(), log, cfg)
	if err != nil {
		log.Error("couldn't run server", sl.Err(err))
		return
	}

	go srv.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	srv.GRPCSrv.Stop()

	log.Info("application stopped")
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
