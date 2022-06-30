package handler

import (
	"context"
	"github.com/nuntiodev/hera/repository/user_repository"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUser creates a user.
func (h *defaultHandler) CreateUser(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		userRepository user_repository.UserRepository
	)
	if req.User == nil {
		return nil, status.Error(codes.InvalidArgument, "missing user")
	}
	userRepository, err = h.repository.UserRepositoryBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	err = userRepository.Create(ctx, mercuryUserToHeraUser(req.User))
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{}, nil
}
