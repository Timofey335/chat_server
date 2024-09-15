package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/Timofey335/platform_common/pkg/db"
	"github.com/fatih/color"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteChat - удаляет чат
func (r *repo) DeleteChat(ctx context.Context, chatId int64) (*emptypb.Empty, error) {
	var id int64

	builderDelete := sq.Delete(tableChats).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: chatId}).
		Suffix("RETURNING id")

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "chat_repository.DeleteChat",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		log.Println(color.HiMagentaString("error while deleting the chat: %v, with ctx: %v", err, ctx))
		return nil, err
	}

	log.Println(color.HiMagentaString("deleted chat: id %d, with ctx: %v", id, ctx))

	return &emptypb.Empty{}, nil
}
