package handler

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/conversations"
)

func (h *defaultHandler) DeleteConversation(ctx context.Context, req *go_mercury.MercuryRequest) (*go_mercury.MercuryResponse, error) {
	var (
		conversationRepository conversations.Conversations
		err                    error
	)
	conversationRepository, err = h.repository.ConversationsBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	err = conversationRepository.Delete(ctx, req.Conversation)
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{}, nil
}
