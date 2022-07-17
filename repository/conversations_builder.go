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

func (b *conversationsBuilder) SetNamespace(namespace string) ConversationsBuilder {
	b.namespace = namespace
	return b
}

func (b *conversationsBuilder) Build(ctx context.Context) (conversations.Conversations, error) {
	if b.namespace == "" {
		b.namespace = "nuntio-db"
	}
	collection := b.client.Database(b.namespace).Collection("mercury_conversations")
	return conversations.New(ctx, collection), nil
}

func (m *mongoRepository) ConversationsBuilder() ConversationsBuilder {
	return &conversationsBuilder{client: m.mongoClient}
}
