package conversations

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *mongoConversations) List(ctx context.Context, user *go_mercury.User, from, to int64) (conversations []*go_mercury.Conversation, err error) {
	var (
		cursor *mongo.Cursor
	)
	if from > to {
		return nil, status.Error(codes.InvalidArgument, "from is greater than to")
	}
	if cursor, err = c.collection.Find(ctx, bson.M{"users": user.Id}, options.Find().SetLimit(from-to).SetSkip(from)); err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &conversations); err != nil {
		return nil, err
	}
	return conversations, nil
}
