package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "go_grpc_chat/pb"
)

const (
	port = ":50051"
)

var user map[string]string

type server struct {
	pb.UnimplementedChatTaskServer
}

func (s *server) Login(ctx context.Context, info *pb.UserInfo) (*pb.LoginResponse, error) {
	log.Printf("User Login: %v",info.UserName)
	user[info.UserName] = "Active"
	return &pb.LoginResponse{Response: pb.ResponseType_SUCCESS}, nil
}

func (s *server) Logout(ctx context.Context, info *pb.UserInfo) (*pb.LogoutResponse, error) {
	log.Printf("User Login: %v",info.UserName)
	delete(user, info.UserName)
	return &pb.LogoutResponse{Response: pb.ResponseType_SUCCESS}, nil
}

func (s * server) Search(ctx context.Context, info *pb.UserInfo) (*pb.UserList, error) {
	log.Printf("User Search: %v",info.UserName)
	return &pb.UserList{UserNameActiveMap: user}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatTaskServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}