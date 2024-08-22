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
	"github.com/Timofey335/chat-server/internal/service"
	chatService "github.com/Timofey335/chat-server/internal/service/chat"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=chats user=user password=userpassword sslmode=disable"
)

type server struct {
	desc.UnimplementedChatServerV1Server
	chatService service.ChatService
}

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

// func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
// 	chatObj, err := s.chatService.CreateChat(ctx, converter.ToChatCreateFromDesc(req))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &desc.CreateChatResponse{
// 		Id: chatObj,
// 	}, nil
// }

// func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
// 	_, err := s.chatService.DeleteChat(ctx, req.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &emptypb.Empty{}, nil
// }

// func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
// 	_, err := s.chatService.SendMessage(ctx, converter.ToSendMessageFromDesc(req))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &emptypb.Empty{}, nil
// }
