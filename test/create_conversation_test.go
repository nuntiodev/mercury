package test

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateConversation(t *testing.T) {
	// create user in client one
	user := &go_mercury.User{FirstName: "Test", LastName: "User"}
	respOne, err := mercuryClientOne.CreateUser(context.Background(), &go_mercury.MercuryRequest{User: user})
	assert.NoError(t, err)
	assert.NotNil(t, respOne)
	assert.NotNil(t, respOne.User)
	// create conversation using client two
	respTwo, err := mercuryClientTwo.CreateConversation(context.Background(), &go_mercury.MercuryRequest{
		Conversation: &go_mercury.Conversation{
			AdminId: respOne.User.Id,
			Users:   []string{respOne.User.Id},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, respTwo)
	assert.NotNil(t, respTwo.Conversation)
	// get conversation using client one
	conversations, err := mercuryClientTwo.ListConversations(context.Background(), &go_mercury.MercuryRequest{
		User: &go_mercury.User{Id: respOne.User.Id},
		From: 0,
		To:   1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, conversations)
	assert.Len(t, conversations.Conversations, 1)
	assert.Equal(t, conversations.Conversations[0], respTwo.Conversation.Id)
}
