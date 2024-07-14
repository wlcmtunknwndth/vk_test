package gRPCApp

import (
	"context"
	"fmt"
	"github.com/wlcmtunknwndth/vk_test/internal/grpc/docsGRPC"
	docsv1 "github.com/wlcmtunknwndth/vk_test/proto/gen/go"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

const scope = "internal.app.grpc."

type DocsHandler interface {
	Create(ctx context.Context, document *docsv1.UserTDocument) (*docsv1.TDocument, error)
	Update(ctx context.Context, document *docsv1.UserTDocument) (*docsv1.TDocument, error)
	Process(ctx context.Context, document *docsv1.TDocument) (*docsv1.TDocument, error)
	Get(ctx context.Context, url string) (*docsv1.TDocument, error)
}

type App struct {
	log         *slog.Logger
	gRPCServer  *grpc.Server
	docsHandler DocsHandler
	port        int
}

func New(log *slog.Logger, port int, handler DocsHandler) *App {
	gRPCServer := grpc.NewServer()

	docsGRPC.Register(gRPCServer, handler, log)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = scope + "Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = scope + "Stop"

	a.log.With(slog.String("op", op)).Info("stooping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
