package langchain_service

import (
	"context"
	pb "langchainGoChatBot/gen/pb-go/langchainGoChatBot/langchain"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedLangchainServiceServer
}

func (s *server) ExampleMethod(ctx context.Context, req *pb.ExampleRequest) (*pb.ExampleResponse, error) {
	response := "Hello, " + req.GetMessage()
	return &pb.ExampleResponse{Response: response}, nil
}

func RegisterGRPCServer(s *grpc.Server) {
	pb.RegisterLangchainServiceServer(s, &server{})
	reflection.Register(s)
}

func Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	RegisterGRPCServer(s)

	log.Println("Starting gRPC server on port :50051...")
	return s.Serve(lis)
}
