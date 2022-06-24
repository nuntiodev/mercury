package conversations

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *mongoConversations) Get(ctx context.Context, conversation *go_mercury.Conversation) (*go_mercury.Conversation, error) {
	if conversation == nil {
		return nil, ConversationIsNil
	} else if conversation.Id == "" {
		return nil, ConversationIdIsEmpty
	}
	var resp go_mercury.Conversation
	if err := c.collection.FindOne(ctx, bson.M{"_id": conversation.Id}).Decode(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
