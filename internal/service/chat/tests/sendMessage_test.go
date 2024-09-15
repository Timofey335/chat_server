package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/Timofey335/platform_common/pkg/db"
	dbMocks "github.com/Timofey335/platform_common/pkg/db/mocks"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Timofey335/chat-server/internal/model"
	"github.com/Timofey335/chat-server/internal/repository"
	repoMocks "github.com/Timofey335/chat-server/internal/repository/mocks"
	"github.com/Timofey335/chat-server/internal/service/chat"
)

func TestSendMessage(t *testing.T) {
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.Message
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id     = gofakeit.Int64()
		userId = gofakeit.Int64()
		chatId = gofakeit.Int64()
		text   = gofakeit.Letter()

		req = &model.Message{
			ID:     id,
			UserId: userId,
			ChatId: chatId,
			Text:   text,
		}

		serviceErr = fmt.Errorf("service error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               *emptypb.Empty
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		txManagerMock      txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.SendMessageMock.Expect(ctx, req).Return(&emptypb.Empty{}, nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.SendMessageMock.Expect(ctx, req).Return(nil, serviceErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			chatRepoMock := tt.chatRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)

			service := chat.NewService(chatRepoMock, txManagerMock)

			resHandler, err := service.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
