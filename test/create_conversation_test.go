package test

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateConversation(t *testing.T) {
	// create users in client one
	firstUser := &go_mercury.User{FirstName: "First", LastName: "User"}
	secondUser := &go_mercury.User{FirstName: "Second", LastName: "User"}
	first, err := mercuryClientOne.CreateUser(context.Background(), &go_mercury.MercuryRequest{User: firstUser})
	assert.NoError(t, err)
	assert.NotNil(t, first)
	assert.NotNil(t, first.User)
	second, err := mercuryClientOne.CreateUser(context.Background(), &go_mercury.MercuryRequest{User: secondUser})
	assert.NoError(t, err)
	assert.NotNil(t, second)
	assert.NotNil(t, second.User)
	// create conversation using client two
	respTwo, err := mercuryClientTwo.CreateConversation(context.Background(), &go_mercury.MercuryRequest{
		Conversation: &go_mercury.Conversation{
			AdminId: first.User.Id,
			Users:   []string{first.User.Id, second.User.Id},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, respTwo)
	assert.NotNil(t, respTwo.Conversation)
	// get conversation using client one
	conversations, err := mercuryClientTwo.ListConversations(context.Background(), &go_mercury.MercuryRequest{
		User: &go_mercury.User{Id: first.User.Id},
		From: 0,
		To:   1,
	})
	assert.NoError(t, err)
	assert.NotNil(t, conversations)
	assert.Len(t, conversations.Conversations, 1)
	assert.Equal(t, conversations.Conversations[0].Id, respTwo.Conversation.Id)
}
