package handler

import (
	"context"
	"github.com/nuntiodev/hera-sdks/go_hera"
	"github.com/nuntiodev/hera/repository/user_repository"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser finds a user by id.
func (h *defaultHandler) GetUser(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		userRepository user_repository.UserRepository
		user          *go_hera.User
	)
	if req.User == nil {
		return nil, status.Error(codes.InvalidArgument, "missing user")
	} else if req.User.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "missing user.id")
	}
	userRepository, err = h.repository.UserRepositoryBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	user, err = userRepository.Get(ctx, &go_hera.User{Id: req.User.Id})
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{User: heraUserToMercuryUser(user)}, nil
}
