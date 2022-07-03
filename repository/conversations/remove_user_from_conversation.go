package conversations

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RemoveUserFromConversation removes a user from a conversation, both given by their id.
func (c *mongoConversations) RemoveUserFromConversation(ctx context.Context, conversation *go_mercury.Conversation, user *go_mercury.User) error {
	if conversation == nil {
		return ConversationIsNil
	} else if conversation.Id == "" {
		return ConversationIdIsEmpty
	} else if user == nil {
		return UserIsNil
	} else if user.Id == "" {
		return UserIdIsEmpty
	}
	result, err := c.collection.UpdateByID(ctx, conversation.Id, bson.M{"$pull": bson.M{"users": user.Id}})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return status.Error(codes.NotFound, "no conversation with id " + conversation.Id)
	}
	return nil
}
