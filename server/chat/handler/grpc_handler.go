package handler

import (
	"context"
	//grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	//grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go_grpc_chat/pb"
	app "go_grpc_chat/server/chat/domain/application"
	"log"
)


func InitGrpcHandler( app *app.Application) *ChatServer{
	return &ChatServer{app: app}
}

type ChatServer struct {
	pb.UnimplementedChatTaskServer
	app *app.Application
}

func (s *ChatServer) SignUp(ctx context.Context, info *pb.UserInfo) (*pb.SignupResponse, error){
	log.Printf("%v signup", info.Username)
	return s.app.AuthApplication.SignUp(info), nil
}

func (s *ChatServer) Login(ctx context.Context, info *pb.UserInfo) (*pb.LoginResponse, error) {
	return s.app.AuthApplication.Login(info), nil
}

func (s *ChatServer) Logout(ctx context.Context, info *pb.UserInfo) (*pb.LogoutResponse, error) {
	return s.app.AuthApplication.Logout(info), nil
}

func (s * ChatServer) Search(ctx context.Context, info *pb.UserInfo) (*pb.UserList, error) {
	return &pb.UserList{UserNameActiveMap: s.app.ChatApplication.Search(info)}, nil
}

func (s * ChatServer) Chat(stream pb.ChatTask_ChatServer) error {
	return s.app.ChatApplication.Chat(stream)
}

