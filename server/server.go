package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	userpb "rakoon/rakoon-reborn/user/pbs"

	"google.golang.org/grpc"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

type userServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

// GetFeature returns the feature at the given point.
func (s *userServiceServer) SayHello(ctx context.Context, point *userpb.Input) (*userpb.Output, error) {
	fmt.Println("HELLO")
	return &userpb.Output{Output: "Hello From the Server!"}, nil
}

func newServer() *userServiceServer {
	// s := &userServiceServer{routeNotes: make(map[string][]*pb.RouteNote)}
	s := &userServiceServer{}
	// s.loadFeatures(*jsonDBFile)
	return s
}

func main() {
	flag.Parse()
    fmt.Println(*port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}