package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Timofey335/chat-server/internal/converter"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	_, err := i.chatService.SendMessage(ctx, converter.ToSendMessageFromDesc(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
