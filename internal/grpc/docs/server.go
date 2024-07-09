package docs

import (
	"context"
	"github.com/wlcmtunknwndth/vk_test/internal/domain/models"
	docsv1 "github.com/wlcmtunknwndth/vk_test/proto/gen/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler interface {
	Create(ctx context.Context, document *models.UserTDocument) (*models.TDocument, error)
	Update(ctx context.Context, document *models.UserTDocument) (*models.TDocument, error)
	Process(ctx context.Context, document *models.TDocument) (*models.TDocument, error)
}

const (
	internalServerError = "internal server error"
)

type serverAPI struct {
	docsv1.UnimplementedDocumentsServer
	docs Handler
}

func (s *serverAPI) Create(ctx context.Context, req *docsv1.UserTDocument) (*docsv1.TDocument, error) {
	tdoc, err := s.docs.Create(ctx, protocUsrDocToUsrDoc(req))
	if err != nil {
		// TODO: add some extra error cases
		return nil, status.Error(codes.Internal, internalServerError)
	}

	return docToProtocDoc(tdoc), nil
}

func (s *serverAPI) Update(ctx context.Context, req *docsv1.UserTDocument) (*docsv1.TDocument, error) {
	tdoc, err := s.docs.Update(ctx, protocUsrDocToUsrDoc(req))
	if err != nil {
		// TODO: handle some other errors
		return nil, status.Error(codes.Internal, internalServerError)
	}

	return docToProtocDoc(tdoc), nil
}

func (s *serverAPI) Process(ctx context.Context, req *docsv1.TDocument) (*docsv1.TDocument, error) {
	tdoc, err := s.docs.Process(ctx, protocDocToDoc(req))
	if err != nil {
		// TODO: add some extra error handlers
		return nil, status.Error(codes.Internal, internalServerError)
	}

	return docToProtocDoc(tdoc), nil
}

func protocDocToDoc(req *docsv1.TDocument) *models.TDocument {
	return &models.TDocument{
		Url:            req.Url,
		PubDate:        req.PubDate,
		FetchTime:      req.FetchTime,
		Text:           req.Text,
		FirstFetchTime: req.FirstFetchTime,
	}
}

func docToProtocDoc(req *models.TDocument) *docsv1.TDocument {
	return &docsv1.TDocument{
		Url:            req.Url,
		PubDate:        req.PubDate,
		FetchTime:      req.FetchTime,
		Text:           req.Text,
		FirstFetchTime: req.FirstFetchTime,
	}
}

func usrDocToProtocUsrDoc(req *models.UserTDocument) *docsv1.UserTDocument {
	return &docsv1.UserTDocument{
		Url:  req.Url,
		Text: req.Text,
	}
}

func protocUsrDocToUsrDoc(req *docsv1.UserTDocument) *models.UserTDocument {
	return &models.UserTDocument{
		Url:  req.Url,
		Text: req.Text,
	}
}
