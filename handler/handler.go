package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository"
	"go.uber.org/zap"
)

type Handler interface {
	Ping(ctx context.Context, req *go_mercury.MercuryRequest) (*go_mercury.MercuryResponse, error)
	Heartbeat(ctx context.Context, req *go_mercury.MercuryRequest) (*go_mercury.MercuryResponse, error)
}

type defaultHandler struct {
	logger     *zap.Logger
	repository *repository.Repository
}

func New(logger *zap.Logger, repository *repository.Repository) (Handler, error) {
	logger.Info("creating handler...")
	return &defaultHandler{
		logger:     logger,
		repository: repository,
	}, nil
}
