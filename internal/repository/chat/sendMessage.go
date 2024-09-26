package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/Timofey335/platform_common/pkg/db"
	"github.com/fatih/color"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Timofey335/chat-server/internal/model"
)

// SendMessage - отправляет сообщение в чат
func (r *repo) SendMessage(ctx context.Context, message *model.Message) (*emptypb.Empty, error) {
	builderInsert := sq.Insert(tableMessages).
		PlaceholderFormat(sq.Dollar).
		Columns(userIdColumn, chatIdColumn, textColumn, createdAtColumn).
		Values(message.UserId, message.ChatId, message.Text, message.CreatedAt)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.SendMessage",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Println(color.HiMagentaString("error while sending the message: %v, with ctx: %v", err, ctx))
		return nil, err
	}

	log.Println(color.BlueString("sent message: %s from user id %s at %s, with ctx: %v", message.Text, message.UserId, message.CreatedAt, ctx))

	return &emptypb.Empty{}, nil
}
