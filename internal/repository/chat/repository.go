package chat

import (
	"github.com/Timofey335/platform_common/pkg/db"

	"github.com/Timofey335/chat-server/internal/repository"
)

const (
	tableChats = "chats"

	idColumn    = "id"
	nameColumn  = "name"
	usersColumn = "users"

	tableMessages = "messages"

	userIdColumn    = "user_id"
	chatIdColumn    = "chat_id"
	textColumn      = "text"
	createdAtColumn = "created_at"
)

type repo struct {
	db db.Client
}

// NewChat - создает новый объект repo
func NewChat(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}
