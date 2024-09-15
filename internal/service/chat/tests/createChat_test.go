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

	"github.com/Timofey335/chat-server/internal/model"
	"github.com/Timofey335/chat-server/internal/repository"
	repoMocks "github.com/Timofey335/chat-server/internal/repository/mocks"
	"github.com/Timofey335/chat-server/internal/service/chat"
)

func TestCreateChat(t *testing.T) {
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.Chat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		usernames = []string{"Bob", "Maria", "John"}

		serviceErr = fmt.Errorf("service error")

		req = &model.Chat{
			ID:    id,
			Name:  name,
			Users: usernames,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		chatRepositoryMock chatRepositoryMockFunc
		txManagerMock      txManagerMockFunc
	}{{
		name: "success case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: id,
		err:  nil,
		chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
			mock := repoMocks.NewChatRepositoryMock(mc)
			mock.CreateChatMock.Expect(ctx, req).Return(id, nil)
			return mock
		},
		txManagerMock: func(mc *minimock.Controller) db.TxManager {
			mock := dbMocks.NewTxManagerMock(mc)
			mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
				return f(ctx)
			})
			return mock
		},
	},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  serviceErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repoMocks.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctx, req).Return(0, serviceErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
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

			resHandler, err := service.CreateChat(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}

}
