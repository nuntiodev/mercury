package repository

import (
	"context"
	"github.com/nuntiodev/mercury/repository/messages"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessagesBuilder interface {
	SetNamespace(namespace string) MessagesBuilder
	Build(ctx context.Context) (messages.Messages, error)
}

type messagesBuilder struct {
	namespace string
	client    *mongo.Client
}

func (b *messagesBuilder) SetNamespace(namespace string) MessagesBuilder {
	b.namespace = namespace
	return b
}

func (b *messagesBuilder) Build(ctx context.Context) (messages.Messages, error) {
	if b.namespace == "" {
		b.namespace = "nuntio-db"
	}
	collection := b.client.Database(b.namespace).Collection("mercury_conversations")
	return messages.New(ctx, collection), nil
}

func (m *mongoRepository) MessagesBuilder() MessagesBuilder {
	return &messagesBuilder{client: m.mongoClient}
}
