package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteChat - удаляет чат
func (s *serv) DeleteChat(ctx context.Context, chatId int64) (*emptypb.Empty, error) {
	_, err := s.chatRepository.DeleteChat(ctx, chatId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
