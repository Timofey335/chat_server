package model

import "time"

type Chat struct {
	ID    int64
	Name  string
	Users []string
}

type Message struct {
	ID        int64
	UserId    int64
	ChatId    int64
	Text      string
	CreatedAt time.Time
}
