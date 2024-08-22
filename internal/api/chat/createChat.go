package chat

import (
	"context"

	"github.com/Timofey335/chat-server/internal/converter"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

func (i *Implementation) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	chatObj, err := i.chatService.CreateChat(ctx, converter.ToChatCreateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateChatResponse{
		Id: chatObj,
	}, nil
}
