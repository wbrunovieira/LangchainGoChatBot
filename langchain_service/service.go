package langchain_service

import (
	"context"
	"log"
	"net"

	pb "langchainGoChatBot/protobuf"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLangchainServiceServer
}

func (s *server) ExampleMethod(ctx context.Context, req *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	return &pb.ExampleResponse{Response: "Hello " + req.Message}, nil
}

func Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLangchainServiceServer(grpcServer, &server{})

	log.Println("Langchain service is running on port 8080")
	return grpcServer.Serve(lis)
}
