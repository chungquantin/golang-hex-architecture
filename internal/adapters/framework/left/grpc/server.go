package grpc

import (
	"log"
	"net"

	"hex/internal/adapters/framework/left/grpc/pb"
	"hex/internal/ports"

	"google.golang.org/grpc"
)

type ServerAdapter struct {
	api ports.APIPort
}

func NewServerAdapter(api ports.APIPort) *ServerAdapter  {
	return &ServerAdapter{api: api}
}

func (grpca ServerAdapter) Run()  {
	var err error

	listen, err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server gRPC server on port 9000: %v", err)
	}
}