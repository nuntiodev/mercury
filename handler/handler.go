package handler

import (
	"errors"
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

var (
	MessageIsNil = errors.New("message is nil")
)

func heraUserModelToMercuryUser(user *models.User) *go_mercury.User {
	return heraUserToMercuryUser(models.UserToProtoUser(user))
}

func heraUserToMercuryUser(user *go_hera.User) *go_mercury.User {
	res := &go_mercury.User{
		Id:        user.GetId(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Image:     user.GetImage(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
	if user.Email != nil {
		res.Email = *user.Email
	}
	return res
}

func mercuryUserToHeraUser(user *go_mercury.User) *go_hera.User {
	return &go_hera.User{
		Id:        user.GetId(),
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
		Image:     &user.Image,
		Email:     &user.Email,
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
}
