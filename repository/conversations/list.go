package conversations

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
)

func (c *mongoConversations) List(ctx context.Context, userId string, from, to int) ([]*go_mercury.Conversation, error) {
	return nil, nil
}
