package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/conversations"
)

// RemoveUserFromConversation removes a user from a conversation, both given by their id.
func (h *defaultHandler) RemoveUserFromConversation(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		conversationRepository conversations.Conversations
	)
	conversationRepository, err = h.repository.ConversationsBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	err = conversationRepository.RemoveUserFromConversation(ctx, req.Conversation, req.User)
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{}, nil
}
