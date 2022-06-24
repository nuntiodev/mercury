package repository

import (
	"context"

	"github.com/nuntiodev/hera/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Repository struct {
	mongoClient *mongo.Client
	ConversationsBuilder
	repository.UserRepositoryBuilder
}

func New(ctx context.Context, mongoClient *mongo.Client, logger *zap.Logger) (*Repository, error) {
	logger.Info("creating repository...")
	return &Repository{
		mongoClient: mongoClient,
	}, nil
}
