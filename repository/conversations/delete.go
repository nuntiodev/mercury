package conversations

import (
	"context"
	"errors"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"go.mongodb.org/mongo-driver/bson"
)

// Delete deletes a conversation by id.
func (c *mongoConversations) Delete(ctx context.Context, conversation *go_mercury.Conversation) error {
	if conversation == nil {
		return ConversationIsNil
	} else if conversation.Id == "" {
		return ConversationIdIsEmpty
	}
	result, err := c.collection.DeleteOne(ctx, bson.M{"_id": conversation.Id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no documents deleted")
	}
	return nil
}
