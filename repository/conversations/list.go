package conversations

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
)

func (c *mongoConversations) List(ctx context.Context, conversation *go_mercury.User, from, to int32) ([]*go_mercury.Conversation, error) {
	return nil, nil
}
