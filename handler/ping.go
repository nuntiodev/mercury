package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
)

func (h *defaultHandler) Ping(ctx context.Context, req *go_mercury.MercuryRequest) (*go_mercury.MercuryResponse, error) {
	return &go_mercury.MercuryResponse{}, nil
}
