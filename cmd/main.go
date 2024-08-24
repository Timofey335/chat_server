package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chatApi "github.com/Timofey335/chat-server/internal/api/chat"
	"github.com/Timofey335/chat-server/internal/config"
	"github.com/Timofey335/chat-server/internal/config/env"
	"github.com/Timofey335/chat-server/internal/repository/chat"
	chatService "github.com/Timofey335/chat-server/internal/service/chat"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	if err := config.Load(configPath); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf(color.RedString("failed listen: %v", err))
	}

	chatRepo := chat.NewChat(pool)
	chatSrv := chatService.NewService(chatRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, chatApi.NewImplementation(chatSrv))
	log.Println(color.BlueString("server listening at %v", lis.Addr()))
	if err := s.Serve(lis); err == nil {
		log.Fatalf(color.RedString("failed to serve: %v", err))
	}
}
