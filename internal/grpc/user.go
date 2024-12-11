package grpc

import (
	"context"
	"time"

	userv1 "github.com/2group/2rent.core-service/pkg/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Api userv1.UserServiceClient
}

func NewClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*Client, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		Api: userv1.NewUserServiceClient(cc),
	}, nil
}
