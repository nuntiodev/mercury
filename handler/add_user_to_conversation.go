package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/conversations"
)

func (h *defaultHandler) AddUserToConversation(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		conversationRepository conversations.Conversations
	)
	conversationRepository, err = h.repository.ConversationsBuilder.SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	err = conversationRepository.AddUserToConversation(ctx, req.Conversation, req.User)
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{}, nil
}
