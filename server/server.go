package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	base "boilerplate/go-grpc/base"
	basepb "boilerplate/go-grpc/base/pbs"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 8080, "The server port")

func main() {
	fmt.Println("Listening on port 8080.")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var baseService base.BaseServiceServer = base.BaseServiceServer{}

	grpcServer := grpc.NewServer()
	basepb.RegisterBaseServiceServer(grpcServer, baseService)
	grpcServer.Serve(lis)
}
