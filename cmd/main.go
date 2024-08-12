package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=chat user=user password=userpassword sslmode=disable"
)

type server struct {
	desc.UnimplementedChatServerV1Server
	pool *pgxpool.Pool
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

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServerV1Server(s, &server{pool: pool})
	log.Println(color.BlueString("server listening at %v", lis.Addr()))
	if err := s.Serve(lis); err == nil {
		log.Fatalf(color.RedString("failed to serve: %v", err))
	}
}

// CreateChat - create a new chat
func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	var chatId int64
	err := s.pool.QueryRow(ctx, `INSERT INTO chats (name, users) VALUES ($1, $2) RETURNING id;`, req.Chatname, req.Usernames).Scan(&chatId)
	if err != nil {
		return nil, err
	}

	log.Println(color.BlueString("create chat: %v, with ctx: %v", req.Chatname, ctx))

	return &desc.CreateChatResponse{
		Id: chatId,
	}, nil
}

// DeleteChat - delete the chat by id
func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	var chatId int64
	err := s.pool.QueryRow(ctx, `DELETE FROM chats WHERE id = $1 RETURNING id;`, req.Id).Scan(&chatId)
	if err != nil {
		log.Println(color.HiMagentaString("error while deleting the chat: %v, with ctx: %v", err, ctx))
		return nil, err
	}

	log.Println(color.HiMagentaString("deleted chat: id %d, with ctx: %v", req.Id, ctx))

	return &emptypb.Empty{}, nil
}

// SendMessage - send message to the server
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	timestamp := req.Timestamp.AsTime().Format(time.UnixDate)
	_, err := s.pool.Exec(ctx, `INSERT INTO messages (user_id, chat_id, text, created_at) VALUES ($1, $2, $3, $4);`, req.FromId, req.ChatId, req.Text, timestamp)
	if err != nil {
		log.Println(color.HiMagentaString("error while sending the message: %v, with ctx: %v", err, ctx))
		return nil, err
	}

	log.Println(color.BlueString("sent message: %s from user id %s at %s, with ctx: %v", req.Text, req.FromId, timestamp, ctx))

	return &emptypb.Empty{}, nil
}
