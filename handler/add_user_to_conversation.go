package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/conversations"
)

func (h *defaultHandler) AddUserToConversation(ctx context.Context, req *go_mercury.MercuryRequest) (*go_mercury.MercuryResponse, error) {
	var (
		conversationRepository conversations.Conversations
		err                    error
	)
	if req.User == nil {
		return nil, UserIsNil
	} else if req.User.Id == "" {
		return nil, UserIdIsEmpty
	} else if req.Conversation == nil {
		return nil, ConversationIsNil
	} else if req.Conversation.Id == "" {
		return nil, ConversationIdIsNil
	}
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
