package repository

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, chat *desc.CreateChatRequest) (*desc.CreateChatResponse, error)
	DeleteChat(ctx context.Context, chatId *desc.DeleteChatRequest) (*emptypb.Empty, error)
	SendMessage(ctx context.Context, message *desc.SendMessageRequest) (*emptypb.Empty, error)
}
