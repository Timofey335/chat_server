package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatServerV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf(color.RedString("failed listen: %v", err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, &server{})
	log.Println(color.BlueString("server listening at %v", lis.Addr()))
	if err := s.Serve(lis); err == nil {
		log.Fatalf(color.RedString("failed to serve: %v", err))
	}
}

// CreateChat - create a new chat
func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	log.Println(color.BlueString("Create chat: %v, with ctx: %v", req.Usernames, ctx))

	return &desc.CreateChatResponse{
		Id: gofakeit.Int64(),
	}, nil
}

// DeleteChat - delete the chat by id
func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Println(color.HiMagentaString("Delete chat: id %d, with ctx: %v", req.Id, ctx))

	return &emptypb.Empty{}, nil
}

// SendMessage - send message to the server
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	timestamp := req.Timestamp.AsTime().Format(time.UnixDate)
	log.Println(color.BlueString("Send message: %s to %s at %s, with ctx: %v", req.Text, req.From, timestamp, ctx))

	return &emptypb.Empty{}, nil
}
