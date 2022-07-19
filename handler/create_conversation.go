package handler

import (
	"context"
	"github.com/nuntiodev/hera-sdks/go_hera"
	"github.com/nuntiodev/hera/repository/user_repository"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/repository/conversations"
	"golang.org/x/sync/errgroup"
	"k8s.io/utils/strings/slices"
)

// CreateConversation creates a new conversation.
// The admin is the user creating the conversation.
func (h *defaultHandler) CreateConversation(ctx context.Context, req *go_mercury.MercuryRequest) (resp *go_mercury.MercuryResponse, err error) {
	var (
		conversationRepository conversations.Conversations
		userRepository         user_repository.UserRepository
		conversation           *go_mercury.Conversation
		errGroup               = errgroup.Group{}
	)
	userRepository, err = h.repository.UserRepositoryBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	// check admin user exists
	errGroup.Go(func() (err error) {
		_, err = userRepository.Get(ctx, &go_hera.User{Id: req.Conversation.AdminId})
		if err != nil {
			return err
		}
		return nil
	})
	// validate that all users in room exists
	errGroup.Go(func() error {
		var users []*go_hera.User
		for _, userId := range req.Conversation.Users {
			users = append(users, &go_hera.User{Id: userId})
		}
		_, err = userRepository.GetMany(ctx, users)
		if err != nil {
			return err
		}
		return nil
	})
	if err = errGroup.Wait(); err != nil {
		return nil, err
	}
	// add admin to room if he has not yet been added
	if !slices.Contains(req.Conversation.Users, req.Conversation.AdminId) {
		req.Conversation.Users = append(req.Conversation.Users, req.Conversation.AdminId)
	}
	// todo: make sure that each id in Conversation.Users is unique (no duplicates)
	conversationRepository, err = h.repository.ConversationsBuilder().SetNamespace(req.Namespace).Build(ctx)
	if err != nil {
		return nil, err
	}
	conversation, err = conversationRepository.Create(ctx, req.Conversation)
	if err != nil {
		return nil, err
	}
	return &go_mercury.MercuryResponse{
		Conversation: conversation,
	}, nil
}
