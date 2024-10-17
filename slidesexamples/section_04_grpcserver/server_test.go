package server

import (
	"almostintegration/pkg/grpcusers"
	"almostintegration/pkg/users"
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcusers.RegisterUserGetServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	clientConn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := grpcusers.NewUserGetClient(clientConn)

	t.Cleanup(func() {
		clientConn.Close()
		s.Stop()
	})

	t.Run("simple call", func(t *testing.T) {
		reply, err := client.Users(context.Background(), &grpcusers.EmptyParams{})
		if err != nil {
			t.Fail()
		}
		if len(reply.Users) != 2 {
			t.Fail()
		}
	})
}

func TestBusinessLogic(t *testing.T) {
	grpcUsers := localUsersToGrpc([]users.User{{"foo", 12}, {"bar", 13}})
	if len(grpcUsers) != 2 {
		t.Fail()
	}
}

func TestImplementation(t *testing.T) {
	s := &server{}
	r, err := s.Users(context.Background(), &grpcusers.EmptyParams{})
	if err != nil {
		t.Fail()
	}
	if len(r.Users) != 2 {
		t.Fail()
	}
}
