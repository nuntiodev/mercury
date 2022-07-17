package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/conversations"
)

// ListConversations lists conversations from req.From (inclusive) to req.To (not inclusive) of a user, given by id.
func (h *defaultHandler) ListConversations(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		conversationRepository conversations.Conversations
		list                   []*go_mercury.Conversation
	)
	conversationRepository, err = h.repository.ConversationsBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	list, err = conversationRepository.List(ctx, req.User, req.From, req.To)
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{Conversations: list}, nil
}
