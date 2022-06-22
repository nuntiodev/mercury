package conversations

import (
	"context"
	"errors"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ConversationIsNil = errors.New("conversation is nil")
)

type Conversations interface {
	Create(ctx context.Context, conversation *go_mercury.Conversation) (*go_mercury.Conversation, error)
	Update(ctx context.Context, conversation *go_mercury.Conversation) (*go_mercury.Conversation, error)
	Get(ctx context.Context, conversation *go_mercury.Conversation) (*go_mercury.Conversation, error)
	List(ctx context.Context, userId string, from, to int) ([]*go_mercury.Conversation, error)
	Delete(ctx context.Context, conversation *go_mercury.Conversation) error
}

type mongoConversations struct {
	collection *mongo.Collection
}

func New(ctx context.Context, collection *mongo.Collection) Conversations {
	return &mongoConversations{
		collection: collection,
	}
}
