package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/conversations"
)

/*
	CreateConversation creates a new conversation with the same adminId as the user creating it.
*/
func (h *defaultHandler) CreateConversation(ctx context.Context, req *go_mercury.MercuryRequest) (*go_mercury.MercuryResponse, error) {
	var (
		c            conversations.Conversations
		conversation *go_mercury.Conversation
		err          error
	)
	c, err = h.repository.ConversationsBuilder.SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	conversation, err = c.Create(ctx, req.Conversation)
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{
		Conversation: conversation,
	}, nil
}
