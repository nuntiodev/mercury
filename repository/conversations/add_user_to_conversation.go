package conversations

import (
	"context"
	"errors"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"go.mongodb.org/mongo-driver/bson"
)

// AddUserToConversation adds a user to a conversation, both given by their id.
func (c *mongoConversations) AddUserToConversation(ctx context.Context, conversation *go_mercury.Conversation, user *go_mercury.User) error {
	if conversation == nil {
		return ConversationIsNil
	} else if conversation.Id == "" {
		return ConversationIdIsEmpty
	} else if user == nil {
		return UserIsNil
	} else if user.Id == "" {
		return UserIdIsEmpty
	}
	//TODO: Check if user is not already added
	result, err := c.collection.UpdateByID(ctx, conversation.Id, bson.M{"$push": bson.M{"users": user.Id}})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no conversation with id " + conversation.Id)
	}
	return nil
}
