package handler

import (
	"context"
	"github.com/nuntiodev/hera-sdks/go_hera"
	"github.com/nuntiodev/hera/models"
	"github.com/nuntiodev/hera/repository/user_repository"
	"github.com/nuntiodev/mercury-proto/go_mercury"
)

// GetAllUsers retrieves all users.
func (h *defaultHandler) GetAllUsers(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		userRepository user_repository.UserRepository
		users          []*models.User
	)
	userRepository, err = h.repository.UserRepositoryBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	users, err = userRepository.List(ctx, &go_hera.Query{})
	if err != nil {
		return nil, err
	}
	res := make([]*go_mercury.User, len(users))
	for i, user := range users {
		res[i] = heraUserModelToMercuryUser(user)
	}
	return &go_mercury.MercuryResponse{Users: res}, nil
}
