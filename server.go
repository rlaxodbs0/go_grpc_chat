package main

import (
	"context"
	"log"
	"net"

	pb "go_grpc_chat/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var userInfo map[string]string
var userStatus map[string]string

type server struct {
	pb.UnimplementedChatTaskServer
}

func (s *server) Signup(ctx context.Context, info *pb.UserInfo) (*pb.SignupResponse, error){
	_, exists := userInfo[info.UserName]
	if exists {
		return &pb.SignupResponse{Response: pb.ResponseType_ALREADYEXISTS}, nil
	}
	userInfo[info.UserName] = info.Password
	log.Printf("User Signup: %v",info.UserName)
	return &pb.SignupResponse{Response:pb.ResponseType_SUCCESS}, nil
}

func (s *server) Login(ctx context.Context, info *pb.UserInfo) (*pb.LoginResponse, error) {
	_, exists := userInfo[info.UserName]
	if !exists {
		return &pb.LoginResponse{Response: pb.ResponseType_NOMATCH}, nil
	}
	log.Printf("User Login: %v",info.UserName)
	userStatus[info.UserName] = "Active"
	return &pb.LoginResponse{Response: pb.ResponseType_SUCCESS}, nil
}

func (s *server) Logout(ctx context.Context, info *pb.UserInfo) (*pb.LogoutResponse, error) {
	log.Printf("User Logout: %v",info.UserName)
	userStatus[info.UserName] = "Inactive"
	return &pb.LogoutResponse{Response: pb.ResponseType_SUCCESS}, nil
}

func (s * server) Search(ctx context.Context, info *pb.UserInfo) (*pb.UserList, error) {
	log.Printf("User Search: %v",info.UserName)
	return &pb.UserList{UserNameActiveMap: userStatus}, nil
}

func main() {
	userInfo = make(map[string]string)
	userStatus = make(map[string]string)
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