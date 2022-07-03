package conversations

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *mongoConversations) Update(ctx context.Context, conversation *go_mercury.Conversation) (*go_mercury.Conversation, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
