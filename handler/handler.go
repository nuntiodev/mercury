package handler

import (
	"errors"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository"
	"go.uber.org/zap"
)

var (
	UserIsNil     = errors.New("user is nil")
	UserIdIsEmpty = errors.New("conversation id is empty")
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
