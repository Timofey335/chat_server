package tests

// func TestCreateChat(t *testing.T) {
// 	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

// 	type args struct {
// 		ctx context.Context
// 		req *desc.CreateChatRequest
// 	}

// 	var (
// 		ctx = context.Background()
// 		mc  = minimock.NewController(t)

// 		id        = gofakeit.Int64()
// 		name      = gofakeit.Name()
// 		usernames = []string{"Bob", "Maria", "John"}

// 		req = &desc.CreateChatRequest{
// 			Chatname:  name,
// 			Usernames: usernames,
// 		}

// 		res = &desc.CreateChatResponse{
// 			Id: id,
// 		}
// 	)

// 	defer t.Cleanup(mc.Finish)

// 	tests := []struct {
// 	}
// }
