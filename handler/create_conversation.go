package handler

import (
	"context"
	"github.com/nuntiodev/hera-proto/go_hera"
	"github.com/nuntiodev/hera/repository/user_repository"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/conversations"
	"golang.org/x/sync/errgroup"
)

/*
	CreateConversation creates a new conversation with the same adminId as the user creating it.
*/
func (h *defaultHandler) CreateConversation(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		c            conversations.Conversations
		u            user_repository.UserRepository
		conversation *go_mercury.Conversation
		errGroup     = errgroup.Group{}
	)
	// check admin user exists
	errGroup.Go(func() (err error) {
		u, err = h.repository.UserRepositoryBuilder.SetNamespace(req.Namespace).Build(ctx)
		if err != nil {
			return err
		}
		u.Get(ctx, &go_hera.User{Id: req.User.GetId(), Email: req.User.Email, Username:})
		return nil
	})
	// check users in conversation exists
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
