package chat

import (
	"github.com/Timofey335/chat-server/internal/service"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

type Implementation struct {
	desc.UnimplementedChatServerV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
