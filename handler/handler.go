package handler

import (
	"github.com/nuntiodev/hera-sdks/go_hera"
	"github.com/nuntiodev/hera/models"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository"
	"go.uber.org/zap"
)

type defaultHandler struct {
	logger     *zap.Logger
	repository repository.Repository
}

func New(logger *zap.Logger, repository repository.Repository) (go_mercury.ServiceServer, error) {
	logger.Info("creating handler...")
	return &defaultHandler{
		logger:     logger,
		repository: repository,
	}, nil
}

func heraUserModelToMercuryUser(user *models.User) *go_mercury.User {
	return heraUserToMercuryUser(models.UserToProtoUser(user))
}

func heraUserToMercuryUser(user *go_hera.User) *go_mercury.User {
	res := &go_mercury.User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	if user.Username != nil {
		res.Name = *user.Username
	}
	if user.Image != nil {
		res.Image = *user.Image
	}
	if user.Email != nil {
		res.Email = *user.Email
	}
	return res
}

func mercuryUserToHeraUser(user *go_mercury.User) *go_hera.User {
	return &go_hera.User{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  &user.Name,
		Image:     &user.Image,
		Email:     &user.Email,
	}
}
