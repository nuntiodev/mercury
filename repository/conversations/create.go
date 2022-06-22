package conversations

import (
	"context"
	"errors"
	"github.com/nuntiodev/mercury-proto/go_mercury"
)

func (c *mongoConversations) Create(ctx context.Context, conversation *go_mercury.Conversation) (*go_mercury.Conversation, error) {
	// validate that the conversation is valid
	if conversation == nil {
		return nil, ConversationIsNil
	} else if conversation.AdminId == "" {
		return nil, errors.New("admin id is empty")
	} else if len(conversation.Users) == 0 {
		return nil, errors.New("no users in conversation")
	}
	if _, err := c.collection.InsertOne(ctx, conversation); err != nil {
		return nil, err
	}
	return conversation, nil
}
