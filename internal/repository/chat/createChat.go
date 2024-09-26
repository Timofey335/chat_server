package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/Timofey335/platform_common/pkg/db"

	"github.com/Timofey335/chat-server/internal/model"
)

// CreateChat - создает новый чат
func (r *repo) CreateChat(ctx context.Context, chat *model.Chat) (int64, error) {
	var chatId int64

	builderInsert := sq.Insert(tableChats).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, usersColumn).
		Values(chat.Name, chat.Users).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.CreateChat",
		QueryRaw: query,
	}

	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatId); err != nil {
		return 0, err
	}

	return chatId, nil
}
