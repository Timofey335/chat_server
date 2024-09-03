package chat

import (
	"github.com/Timofey335/chat-server/internal/service"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

// Implementation - структура содержащая заглушки GRPC методов (в случае если они еще не созданы) и
// объект (интерфейс) сервисного слоя
type Implementation struct {
	desc.UnimplementedChatServerV1Server
	chatService service.ChatService
}

// NewImplementation - конструктор, который возвращает объект сервисного слоя
func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
