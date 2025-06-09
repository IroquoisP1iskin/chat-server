package main

import (
	"context"
	"fmt"
	"github.com/IroquoisP1iskin/chat-server/pkg/chat_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"

	desc "github.com/IroquoisP1iskin/auth/pkg/note_v1"
)

const (
	address = "localhost:50051"
	noteID  = 23
)

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
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := desc.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Fatalf("could not get note: %v", err)
	} else {
		log.Printf("Note: %v", r)
	}

	log.Printf("%s%s", color.GreenString("Note:"), color.GreenString(" %v", r.GetEmail()))
}
