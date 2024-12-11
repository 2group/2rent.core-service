package grpc

import (
	"context"
	"time"

	operationv1 "github.com/2group/2rent.core-service/pkg/gen/go/operation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OperationClient struct {
	Api operationv1.OperationServiceClient
}

func NewOperationClient(ctx context.Context, addr string, timeout time.Duration, retriesCount int) (*OperationClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &OperationClient{
		Api: operationv1.NewOperationServiceClient(cc),
	}, nil
}
