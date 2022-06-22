package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Repository struct {
	mongoClient *mongo.Client
	ConversationsBuilder
}

func New(ctx context.Context, mongoClient *mongo.Client, logger *zap.Logger) (*Repository, error) {
	logger.Info("creating repository...")
	return &Repository{
		mongoClient: mongoClient,
	}, nil
}
