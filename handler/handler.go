package handler

import (
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository"
	"go.uber.org/zap"
)

type defaultHandler struct {
	logger     *zap.Logger
	repository *repository.Repository
}

func New(logger *zap.Logger, repository *repository.Repository) (go_mercury.ServiceServer, error) {
	logger.Info("creating handler...")
	return &defaultHandler{
		logger:     logger,
		repository: repository,
	}, nil
}
