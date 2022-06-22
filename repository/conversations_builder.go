package repository

import (
	"context"
	"github.com/nuntiodev/mercury/repository/conversations"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConversationsBuilder interface {
	SetNamespace(namespace string) ConversationsBuilder
	Build(ctx context.Context) (conversations.Conversations, error)
}

type conversationsBuilder struct {
	namespace string
	client    *mongo.Client
}

func (cb *conversationsBuilder) SetNamespace(namespace string) ConversationsBuilder {
	cb.namespace = namespace
	return cb
}

func (cb *conversationsBuilder) Build(ctx context.Context) (conversations.Conversations, error) {
	if cb.namespace == "" {
		cb.namespace = "nuntio-db"
	}
	collection := cb.client.Database(cb.namespace).Collection("mercury_conversations")
	return conversations.New(ctx, collection), nil
}
