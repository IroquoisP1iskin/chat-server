package main

import (
	"context"
	"fmt"
	"github.com/IroquoisP1iskin/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const address = "127.0.0.1:50001"

type server struct {
	chat_v1.UnimplementedChatServiceServer
}

func (s *server) Create(_ context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	fmt.Printf("#%v", req)
	return &chat_v1.CreateResponse{Id: 11}, nil
}

func (s *server) Delete(_ context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Printf("#%v", req)
	return &emptypb.Empty{}, nil
}

func (s *server) Send(_ context.Context, req *chat_v1.SendMessageRequest) (*emptypb.Empty, error) {
	fmt.Printf("#%v", req)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to create listener: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	chat_v1.RegisterChatServiceServer(grpcServer, &server{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}
}
