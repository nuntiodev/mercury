package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *defaultHandler) AddUserToConversation(ctx context.Context, req *go_mercury.MercuryRequest) (*go_mercury.MercuryResponse, error) {
	//TODO: Implement
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
