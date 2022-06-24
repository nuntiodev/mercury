package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/conversations"
)

func (h *defaultHandler) ListConversations(ctx context.Context, req *go_mercury.MercuryRequest) (*go_mercury.MercuryResponse, error) {
	var (
		conversationRepository conversations.Conversations
		list                   []*go_mercury.Conversation
		err                    error
	)
	if req.User == nil {
		return nil, UserIsNil
	} else if req.User.Id == "" {
		return nil, UserIdIsEmpty
	}
	conversationRepository, err = h.repository.ConversationsBuilder.SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	list, err = conversationRepository.List(ctx, req.User.Id, int(req.From), int(req.To))
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{Conversations: list}, nil
}
