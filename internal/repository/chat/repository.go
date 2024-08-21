package chat

import (
	"context"
	"log"
	"time"

	"github.com/Timofey335/chat-server/internal/repository"
	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

type repo struct {
	db *pgxpool.Pool
}

func NewChat(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}

// CreateChat - create a new chat
func (r *repo) CreateChat(ctx context.Context, chat *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	var chatId int64

	err := r.db.QueryRow(ctx, `INSERT INTO chats (name, users) VALUES ($1, $2) RETURNING id;`, chat.Chatname, chat.Usernames).Scan(&chatId)
	if err != nil {
		return nil, err
	}

	log.Println(color.BlueString("create chat: %v, with ctx: %v", chat.Chatname, ctx))

	return &desc.CreateChatResponse{
		Id: chatId,
	}, nil
}

// DeleteChat - delete the chat by id
func (r *repo) DeleteChat(ctx context.Context, chatId *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	var id int64

	err := r.db.QueryRow(ctx, `DELETE FROM chats WHERE id = $1 RETURNING id;`, chatId.Id).Scan(&id)
	if err != nil {
		log.Println(color.HiMagentaString("error while deleting the chat: %v, with ctx: %v", err, ctx))
		return nil, err
	}

	log.Println(color.HiMagentaString("deleted chat: id %d, with ctx: %v", chatId.Id, ctx))

	return &emptypb.Empty{}, nil
}

// SendMessage - send message to the server
func (r *repo) SendMessage(ctx context.Context, message *desc.SendMessageRequest) (*emptypb.Empty, error) {
	timestamp := message.Timestamp.AsTime().Format(time.UnixDate)
	_, err := r.db.Exec(ctx, `INSERT INTO messages (user_id, chat_id, text, created_at) VALUES ($1, $2, $3, $4);`, message.FromId, message.ChatId, message.Text, timestamp)
	if err != nil {
		log.Println(color.HiMagentaString("error while sending the message: %v, with ctx: %v", err, ctx))
		return nil, err
	}

	log.Println(color.BlueString("sent message: %s from user id %s at %s, with ctx: %v", message.Text, message.FromId, timestamp, ctx))

	return &emptypb.Empty{}, nil
}
