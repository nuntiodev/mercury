package server

import (
	"context"
	"fmt"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/nuntiodev/mercury/handler"
	"github.com/nuntiodev/mercury/repository"
	database "github.com/nuntiodev/x/repositoryx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

var (
	port = ""
)

type Server interface {
	Run() error
}

type defaultServer struct {
	handler handler.Handler
	logger  *zap.Logger
}

func initializeServer() error {
	var ok bool
	port, ok = os.LookupEnv("PORT")
	if !ok || port == "" {
		port = "9000"
	}
	return nil
}

func New(ctx context.Context, logger *zap.Logger) (Server, error) {
	if err := initializeServer(); err != nil {
		return nil, err
	}
	myDatabase, err := database.CreateDatabase(logger)
	if err != nil {
		return nil, err
	}
	mongoClient, err := myDatabase.CreateMongoClient(ctx)
	if err != nil {
		return nil, err
	}
	// create repository
	myRepository, err := repository.New(ctx, mongoClient, logger)
	if err != nil {
		return nil, err
	}
	// create handler
	myHandler, err := handler.New(logger, myRepository)
	if err != nil {
		return nil, err
	}
	return &defaultServer{
		handler: myHandler,
	}, nil
}

func (s *defaultServer) Run() error {
	s.logger.Info("running grpc server...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			s.logger.Fatal(err.Error())
		}
	}(lis)
	s.logger.Info(fmt.Sprintf("server running on port: %s", port))
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	go_mercury.RegisterServiceServer(grpcServer, s.handler)
	return grpcServer.Serve(lis)
}
