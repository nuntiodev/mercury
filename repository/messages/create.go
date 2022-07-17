package messages

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
)

// Create creates a message.
func (c *mongoMessages) Create(ctx context.Context, message *go_mercury.Message) (*go_mercury.Message, error) {
	// validate that the message is valid
	if message == nil {
		return nil, MessageIsNil
	} else if message.Id == "" {
		return nil, MessageIdIsEmpty
	}
	if _, err := c.collection.InsertOne(ctx, message); err != nil {
		return nil, err
	}
	return message, nil
}
