package base

import (
	basepb "boilerplate/go-grpc/base/pbs"
	"context"
)

type BaseServiceServer struct {
	basepb.UnimplementedBaseServiceServer
}

// GetFeature returns the feature at the given point.
func (s BaseServiceServer) Ping(ctx context.Context, point *basepb.Input) (*basepb.Output, error) {
	return &basepb.Output{Output: "Hello From the Server!"}, nil
}
