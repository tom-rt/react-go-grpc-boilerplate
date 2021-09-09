package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	basepb "boilerplate/go-grpc/base/pbs"

	"google.golang.org/grpc"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

type baseServiceServer struct {
	basepb.UnimplementedBaseServiceServer
}

// GetFeature returns the feature at the given point.
func (s *baseServiceServer) SayHello(ctx context.Context, point *basepb.Input) (*basepb.Output, error) {
	fmt.Println("HELLO")
	return &basepb.Output{Output: "Hello From the Server!"}, nil
}

func newServer() *baseServiceServer {
	// s := &userServiceServer{routeNotes: make(map[string][]*pb.RouteNote)}
	s := &baseServiceServer{}
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
	basepb.RegisterBaseServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}