package handler

import (
	"context"
	//grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	//grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go_grpc_chat/pb"
	app "go_grpc_chat/server/chat/domain/application"
)

/*
	rpc CreateUser(CreateUserRequest) returns (CreateUserResponse) {}
    rpc Login(LoginRequest) returns (LoginResponse) {}
    rpc Logout(LogoutRequest) returns (LogoutResponse) {}
    rpc AddFriend(AddFriendRequest) returns (AddFriendResponse) {}
    rpc RemoveFriend(RemoveFriendRequest) returns (RemoveFriendResponse) {}
    rpc GetFriends(GetFriendsRequest) returns (GetFriendsResponse) {}
    rpc Chat(stream Message) returns (stream Message) {}
 */


type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

func (s *ChatServer) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error){
	return app.UserApplication().CreateUser(ctx, request)
}

func (s *ChatServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	return app.UserApplication().Login(ctx, request)
}

func (s *ChatServer) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return app.UserApplication().Logout(ctx, request)
}

func (s * ChatServer) AddFriend(ctx context.Context, request *pb.AddFriendRequest) (*pb.AddFriendResponse, error) {
	return app.UserApplication().AddFriend(ctx, request)
}

func (s * ChatServer) RemoveFriend(ctx context.Context, request *pb.RemoveFriendRequest) (*pb.RemoveFriendResponse, error) {
	return app.UserApplication().RemoveFriend(ctx, request)
}

func (s * ChatServer) GetFriends(ctx context.Context, request *pb.GetFriendsRequest) (*pb.GetFriendsResponse, error) {
	return app.UserApplication().GetFriends(ctx, request)
}

func (s * ChatServer) Chat(stream pb.ChatService_ChatServer) error {
	return app.ChatApplication().Chat(stream)
}

