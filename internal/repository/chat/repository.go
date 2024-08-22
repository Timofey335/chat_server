package chat

import (
	"context"
	"log"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Timofey335/chat-server/internal/model"
	"github.com/Timofey335/chat-server/internal/repository"
)

type repo struct {
	db *pgxpool.Pool
}

func NewChat(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}

// CreateChat - create a new chat
func (r *repo) CreateChat(ctx context.Context, chat *model.Chat) (int64, error) {
	var chatId int64

	err := r.db.QueryRow(ctx, `INSERT INTO chats (name, users) VALUES ($1, $2) RETURNING id;`, chat.Name, chat.Users).Scan(&chatId)
	if err != nil {
		return 0, err
	}

	log.Println(color.BlueString("create chat: %v, with ctx: %v", chat.Name, ctx))

	return chatId, nil
}

// DeleteChat - delete the chat by id
func (r *repo) DeleteChat(ctx context.Context, chatId int64) (*emptypb.Empty, error) {
	var id int64

	err := r.db.QueryRow(ctx, `DELETE FROM chats WHERE id = $1 RETURNING id;`, chatId).Scan(&id)
	if err != nil {
		log.Println(color.HiMagentaString("error while deleting the chat: %v, with ctx: %v", err, ctx))
		return nil, err
	}

	log.Println(color.HiMagentaString("deleted chat: id %d, with ctx: %v", chatId, ctx))

	return &emptypb.Empty{}, nil
}

// SendMessage - send message to the server
func (r *repo) SendMessage(ctx context.Context, message *model.Message) (*emptypb.Empty, error) {
	_, err := r.db.Exec(ctx, `INSERT INTO messages (user_id, chat_id, text, created_at) VALUES ($1, $2, $3, $4);`, message.UserId, message.ChatId, message.Text, message.CreatedAt)
	if err != nil {
		log.Println(color.HiMagentaString("error while sending the message: %v, with ctx: %v", err, ctx))
		return nil, err
	}

	log.Println(color.BlueString("sent message: %s from user id %s at %s, with ctx: %v", message.Text, message.UserId, message.CreatedAt, ctx))

	return &emptypb.Empty{}, nil
}
