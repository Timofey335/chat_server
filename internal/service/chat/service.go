package chat

import (
	"github.com/Timofey335/chat-server/internal/client/db"
	"github.com/Timofey335/chat-server/internal/repository"
	def "github.com/Timofey335/chat-server/internal/service"
)

var _ def.ChatService = (*serv)(nil)

type serv struct {
	chatRepository repository.ChatRepository
	txManager      db.TxManager
}

// NewService - создает новый экземпляр serv
func NewService(chatRepository repository.ChatRepository, txManager db.TxManager) *serv {
	return &serv{
		chatRepository: chatRepository,
		txManager:      txManager,
	}
}
