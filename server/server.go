package main

import (
	"fmt"
	"log"
	"net"

	base "boilerplate/go-grpc/base"
	basepb "boilerplate/go-grpc/base/pbs"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Listening on port 8080.")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var baseService base.BaseServiceServer = base.BaseServiceServer{}

	grpcServer := grpc.NewServer()
	basepb.RegisterBaseServiceServer(grpcServer, baseService)
	grpcServer.Serve(lis)
}
