package converter

import (
	"github.com/Timofey335/chat-server/internal/model"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

func ToChatCreateFromDesc(chat *desc.CreateChatRequest) *model.Chat {
	return &model.Chat{
		Name:  chat.Chatname,
		Users: chat.Usernames,
	}
}

func ToSendMessageFromDesc(message *desc.SendMessageRequest) *model.Message {
	return &model.Message{
		UserId:    message.FromId,
		ChatId:    message.ChatId,
		Text:      message.Text,
		CreatedAt: message.Timestamp.AsTime(),
	}
}
