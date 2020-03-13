package handler

import (
	//grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	//grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go_grpc_chat/pb"
	app "go_grpc_chat/server/chat/domain/application"
	"google.golang.org/grpc"
	"log"
	"net"
	"context"
)



func InitGrpcHandler(port string, app *app.Application) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatTaskServer(s, &ChatServer{app: app})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type ChatServer struct {
	pb.UnimplementedChatTaskServer
	app *app.Application
}

func (s *ChatServer) SignUp(ctx context.Context, info *pb.UserInfo) (*pb.SignupResponse, error){
	log.Printf("%v signup", info.Username)
	return s.app.AuthApplictaion.SignUp(info), nil
}

func (s *ChatServer) Login(ctx context.Context, info *pb.UserInfo) (*pb.LoginResponse, error) {
	return s.app.AuthApplictaion.Login(info), nil
}

func (s *ChatServer) Logout(ctx context.Context, info *pb.UserInfo) (*pb.LogoutResponse, error) {
	return s.app.AuthApplictaion.Logout(info), nil
}

func (s * ChatServer) Search(ctx context.Context, info *pb.UserInfo) (*pb.UserList, error) {
	return &pb.UserList{UserNameActiveMap: s.app.Search(info)}, nil
}

func (s * ChatServer) Chat(stream pb.ChatTask_ChatServer) error {
	return s.app.ChatApplictaion.Chat(stream)
}

