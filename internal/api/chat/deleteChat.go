package chat

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

func (i *Implementation) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	_, err := i.chatService.DeleteChat(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
