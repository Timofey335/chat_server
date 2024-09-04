package model

import (
	"time"
)

// Chat - модель Chat
type Chat struct {
	ID    int64
	Name  string
	Users []string
}

// Message - модель Message
type Message struct {
	ID        int64
	UserId    int64
	ChatId    int64
	Text      string
	CreatedAt time.Time
}
