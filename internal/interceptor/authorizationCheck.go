package interceptor

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	descAccess "github.com/Timofey335/chat-server/pkg/access_v1"
)

var accessToken = flag.String("a", "", "access token")

const servicePort = 50051

// AuthorizarionCheck - интерцептор, проверяет пользователя к эндпоинту
func AuthorizarionCheck(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	flag.Parse()

	creds, err := credentials.NewClientTLSFromFile("cert/service.pem", "")
	if err != nil {
		log.Fatalf("could not process the credentials: %v", err)
	}

	md := metadata.New(map[string]string{"Authorization": "Bearer " + *accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", servicePort),
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v", err)
	}

	cl := descAccess.NewAccessV1Client(conn)

	_, err = cl.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: info.FullMethod,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("Access granted")

	return handler(ctx, req)
}
