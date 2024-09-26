package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Timofey335/chat-server/internal/api/chat"
	"github.com/Timofey335/chat-server/internal/model"
	"github.com/Timofey335/chat-server/internal/service"
	serviceMocks "github.com/Timofey335/chat-server/internal/service/mocks"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

func TestSendMessage(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		fromId    = gofakeit.Int64()
		chatId    = gofakeit.Int64()
		text      = gofakeit.Question()
		timestamp = timestamppb.Now()

		req = &desc.SendMessageRequest{
			FromId:    fromId,
			ChatId:    chatId,
			Text:      text,
			Timestamp: timestamp,
		}

		messageModel = &model.Message{
			UserId:    fromId,
			ChatId:    chatId,
			Text:      text,
			CreatedAt: timestamp.AsTime(),
		}

		serviceErr = fmt.Errorf("service error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, messageModel).Return(&emptypb.Empty{}, nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, messageModel).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)
			api := chat.NewImplementation(userServiceMock)

			resHandler, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})

	}
}
