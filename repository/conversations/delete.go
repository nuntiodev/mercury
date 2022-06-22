package conversations

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
)

func (c *mongoConversations) Delete(ctx context.Context, conversation *go_mercury.Conversation) error {
	return nil
}
