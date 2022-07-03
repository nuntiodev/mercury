package repository

import (
	"context"
	"github.com/nuntiodev/hera/repository"
	hera_repository "github.com/nuntiodev/hera/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Repository interface {
	UserRepositoryBuilder() hera_repository.UserRepositoryBuilder
	ConversationsBuilder() ConversationsBuilder
}

type mongoRepository struct {
	mongoClient           *mongo.Client
	userRepositoryBuilder repository.UserRepositoryBuilder
}

func (m *mongoRepository) UserRepositoryBuilder() hera_repository.UserRepositoryBuilder {
	return m.userRepositoryBuilder
}

func New(ctx context.Context, mongoClient *mongo.Client, logger *zap.Logger) (Repository, error) {
	logger.Info("creating repository...")

	return &mongoRepository{
		mongoClient:           mongoClient,
		userRepositoryBuilder: hera_repository.NewUserRepositoryBuilder(mongoClient),
	}, nil
}
