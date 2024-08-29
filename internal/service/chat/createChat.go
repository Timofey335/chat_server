package chat

import (
	"context"
	"log"

	"github.com/fatih/color"

	"github.com/Timofey335/chat-server/internal/model"
)

func (s *serv) CreateChat(ctx context.Context, chat *model.Chat) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.CreateChat(ctx, chat)
		if errTx != nil {
			return errTx
		}

		log.Println(color.BlueString("created chat: %v, with ctx: %v", chat.Name, ctx))

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
