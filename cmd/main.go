package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chatApi "github.com/Timofey335/chat-server/internal/api/chat"
	"github.com/Timofey335/chat-server/internal/repository/chat"
	chatService "github.com/Timofey335/chat-server/internal/service/chat"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=chats user=user password=userpassword sslmode=disable"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
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
