package docsGRPC

import (
	"context"
	"github.com/wlcmtunknwndth/vk_test/internal/lib/sl"
	docsv1 "github.com/wlcmtunknwndth/vk_test/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type Handler interface {
	Create(ctx context.Context, document *docsv1.UserTDocument) (*docsv1.TDocument, error)
	Update(ctx context.Context, document *docsv1.UserTDocument) (*docsv1.TDocument, error)
	Process(ctx context.Context, document *docsv1.TDocument) (*docsv1.TDocument, error)
	Get(ctx context.Context, url string) (*docsv1.TDocument, error)
}

const (
	internalServerError = "internal server error"
	notFound            = "document not found"
)

type serverAPI struct {
	docsv1.UnimplementedDocumentsServer
	docs Handler
	log  *slog.Logger
}

func Register(grpc *grpc.Server, handler Handler, log *slog.Logger) {
	docsv1.RegisterDocumentsServer(grpc, &serverAPI{
		UnimplementedDocumentsServer: docsv1.UnimplementedDocumentsServer{},
		docs:                         handler,
		log:                          log,
	})
}

func (s *serverAPI) Create(ctx context.Context, req *docsv1.UserTDocument) (*docsv1.TDocument, error) {
	tdoc, err := s.docs.Create(ctx, req)
	if err != nil {
		// TODO: add some extra error cases
		s.log.Error("couldn't create document", sl.Err(err))
		return nil, status.Error(codes.Internal, internalServerError)
	}

	return tdoc, nil
}

func (s *serverAPI) Update(ctx context.Context, req *docsv1.UserTDocument) (*docsv1.TDocument, error) {
	tdoc, err := s.docs.Update(ctx, req)
	if err != nil {
		// TODO: handle some other errors

		s.log.Error("couldn't update document", sl.Err(err))
		return nil, status.Error(codes.Internal, internalServerError)
	}

	return tdoc, nil
}

func (s *serverAPI) Process(ctx context.Context, req *docsv1.TDocument) (*docsv1.TDocument, error) {
	tdoc, err := s.docs.Process(ctx, req)
	if err != nil {
		s.log.Error("couldn't process document", sl.Err(err))
		return nil, status.Error(codes.Internal, internalServerError)
	}

	return tdoc, nil
}

func (s *serverAPI) Get(ctx context.Context, req *docsv1.GetRequest) (*docsv1.TDocument, error) {
	tdoc, err := s.docs.Get(ctx, req.Url)
	if err != nil {

		s.log.Error("couldn't get document", sl.Err(err))
		return nil, status.Error(codes.NotFound, notFound)
	}

	return tdoc, err
}
