package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/messages"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Send creates a message in a conversation by a user, both given by their id.
func (h *defaultHandler) Send(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		message            *go_mercury.Message
		messagesRepository messages.Messages
	)

	if req.Message == nil {
		return nil, MessageIsNil
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	messagesRepository, err = h.repository.MessagesBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	message, err = messagesRepository.Create(ctx, &go_mercury.Message{
		Id:             id.String(),
		ConversationId: req.Message.ConversationId,
		UserId:         req.Message.UserId,
		SentAt:         timestamppb.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{Message: message}, nil
}
