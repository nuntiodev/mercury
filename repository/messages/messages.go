package messages

import (
	"context"
	"errors"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MessageIsNil     = errors.New("message is nil")
	MessageIdIsEmpty = errors.New("message id is empty")
)

type Messages interface {
	Create(ctx context.Context, conversation *go_mercury.Message) (*go_mercury.Message, error)
}

type mongoMessages struct {
	collection *mongo.Collection
}

func New(ctx context.Context, collection *mongo.Collection) Messages {
	return &mongoMessages{
		collection: collection,
	}
}
