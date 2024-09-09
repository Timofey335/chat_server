package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Timofey335/chat-server/internal/model"
)

// SendMessage - отправляет сообщение в чат
func (s *serv) SendMessage(ctx context.Context, message *model.Message) (*emptypb.Empty, error) {
	_, err := s.chatRepository.SendMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
