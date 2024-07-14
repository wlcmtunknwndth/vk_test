package app

import (
	"context"
	"fmt"
	"github.com/wlcmtunknwndth/vk_test/internal/app/gRPCApp"
	"github.com/wlcmtunknwndth/vk_test/internal/broker/nats"
	"github.com/wlcmtunknwndth/vk_test/internal/cache/Redis"
	"github.com/wlcmtunknwndth/vk_test/internal/config"
	"github.com/wlcmtunknwndth/vk_test/internal/lib/sl"
	"github.com/wlcmtunknwndth/vk_test/internal/services/docs"
	"log/slog"
)

const scope = "internal.App."

type App struct {
	GRPCSrv *gRPCApp.App
}

func New(ctx context.Context, log *slog.Logger, cfg *config.Config) (*App, error) {
	const op = scope + "New"
	rds, err := Redis.New(cfg.Redis.Url, cfg.Redis.Password, cfg.Redis.DbOpt)
	if err != nil {
		log.Error("couldn't connect to Redis", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	ns, err := nats.New(log, rds,
		cfg.Nats.Address, cfg.Nats.Retry,
		cfg.Nats.MaxReconnects, cfg.Nats.ReconnectWait,
	)
	if err != nil {
		log.Error("couldn't connect to nats", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	go func(ctx context.Context) {
		if err := ns.Start(ctx); err != nil {
			log.Error("couldn't start Receiver Queue", sl.Op(op), sl.Err(err))
		}
		return
	}(ctx)

	srvc := docs.New(ns, rds)

	return &App{GRPCSrv: gRPCApp.New(log, cfg.GRPC.Port, srvc)}, nil
}
