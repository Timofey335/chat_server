package app

import (
	"context"
	"log"

	"github.com/Timofey335/chat-server/internal/api/chat"
	"github.com/Timofey335/chat-server/internal/client/db"
	"github.com/Timofey335/chat-server/internal/client/db/pg"
	"github.com/Timofey335/chat-server/internal/closer"
	"github.com/Timofey335/chat-server/internal/config"
	"github.com/Timofey335/chat-server/internal/config/env"
	"github.com/Timofey335/chat-server/internal/repository"
	chatRepository "github.com/Timofey335/chat-server/internal/repository/chat"
	"github.com/Timofey335/chat-server/internal/service"
	chatService "github.com/Timofey335/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	chatRepository repository.ChatRepository

	chatService service.ChatService

	servImplementation *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient

}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewChat(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.ChatRepository(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) ServImplementation(ctx context.Context) *chat.Implementation {
	if s.servImplementation == nil {
		s.servImplementation = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.servImplementation
}
