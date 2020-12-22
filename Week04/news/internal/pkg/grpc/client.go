package grpc

import (
	"context"
	"github.com/belson77/Go-001/Week04/news/config"
	rpc "google.golang.org/grpc"
)

func NewClient(ctx context.Context, cf config.Config) (*rpc.ClientConn, error) {
	conn, err := rpc.DialContext(ctx, cf.RPC.Address, rpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
