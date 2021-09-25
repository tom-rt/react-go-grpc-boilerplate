package base

import (
	basepb "boilerplate/go-grpc/base/pbs"
	"context"
)

type BaseServiceServer struct {
	basepb.UnimplementedBaseServiceServer
}

// Basic ping function
func (s BaseServiceServer) Ping(ctx context.Context, input *basepb.Input) (*basepb.Output, error) {
	return &basepb.Output{Output: "Pong"}, nil
}
