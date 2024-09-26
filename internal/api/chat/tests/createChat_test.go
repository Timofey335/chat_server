package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/Timofey335/chat-server/internal/api/chat"
	"github.com/Timofey335/chat-server/internal/model"
	"github.com/Timofey335/chat-server/internal/service"
	serviceMocks "github.com/Timofey335/chat-server/internal/service/mocks"
	desc "github.com/Timofey335/chat-server/pkg/chat_server_v1"
)

func TestCreateChat(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		usernames = []string{"Bob", "Maria", "John"}

		req = &desc.CreateChatRequest{
			Chatname:  name,
			Usernames: usernames,
		}

		chatModel = &model.Chat{
			Name:  name,
			Users: usernames,
		}

		res = &desc.CreateChatResponse{
			Id: id,
		}

		serviceErr = fmt.Errorf("service error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateChatResponse
		err             error
		userServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateChatMock.Expect(ctx, chatModel).Return(id, nil)
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
				mock.CreateChatMock.Expect(ctx, chatModel).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatServiceMock := tt.userServiceMock(mc)
			api := chat.NewImplementation(chatServiceMock)

			resHandler, err := api.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
