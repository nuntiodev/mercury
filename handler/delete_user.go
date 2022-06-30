package handler

import (
	"context"
	"github.com/nuntiodev/hera-sdks/go_hera"
	"github.com/nuntiodev/hera/models"
	"github.com/nuntiodev/hera/repository/user_repository"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteUser deletes a user by id.
func (h *defaultHandler) DeleteUser(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		userRepository user_repository.UserRepository
		user          *models.User
	)
	if req.User == nil {
		return nil, status.Error(codes.InvalidArgument, "missing user")
	}
	userRepository, err = h.repository.UserRepositoryBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	if err = userRepository.Delete(ctx, &go_hera.User{Id: req.User.Id}); err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{User: heraUserModelToMercuryUser(user)}, nil
}
