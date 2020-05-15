package application

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/domain/model"
	"go_grpc_chat/server/chat/domain/repository"
	"go_grpc_chat/server/chat/infrastructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

type userApplication struct{
	repository repository.UserRepositoryImpl
}

var (
	once sync.Once
	userApplicationInstance *userApplication
)

func UserApplication() *userApplication{
	once.Do(func() {
		userApplicationInstance = &userApplication{repository:infrastructure.NewUserRepository()}
	})
	return userApplicationInstance
}

func (s *userApplication) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := s.repository.CreateUser(&model.User{Username:request.UserInfo.Username, Password:request.UserInfo.Password}); err != nil {
		return &pb.CreateUserResponse{Response:pb.ResponseType_FAIL}, err
	}
	return &pb.CreateUserResponse{Response:pb.ResponseType_SUCCESS}, nil
}

func (s *userApplication) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	usr, err := s.repository.GetUserByUsername(request.UserInfo.Username)
	if err != nil {
		errorStatus := status.New(codes.AlreadyExists, fmt.Sprintf("User %s not exists", request.UserInfo.Username))
		return &pb.LoginResponse{Response:pb.ResponseType_FAIL}, errorStatus.Err()
	}
	if usr.Password != request.UserInfo.Password {
		errorStatus := status.New(codes.NotFound, fmt.Sprintf("Password not matched"))
		return &pb.LoginResponse{Response:pb.ResponseType_FAIL}, errorStatus.Err()
	}
	usr.Login()
	if err := s.repository.UpdateUser(usr); err != nil {
		errorStatus := status.New(codes.Unknown, fmt.Sprintf("Login status error"))
		return &pb.LoginResponse{Response:pb.ResponseType_FAIL}, errorStatus.Err()
	}
	return &pb.LoginResponse{Response:pb.ResponseType_SUCCESS}, nil
}

func (s *userApplication) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	usr, err := s.repository.GetUserByUsername(request.UserInfo.Username)
	if err != nil {
		errorStatus := status.New(codes.NotFound, fmt.Sprintf("User %s not exists", request.UserInfo.Username))
		return &pb.LogoutResponse{Response:pb.ResponseType_FAIL}, errorStatus.Err()
	}
	usr.Logout()
	if err := s.repository.UpdateUser(usr); err != nil {
		errorStatus := status.New(codes.Unknown, fmt.Sprintf("Login status error"))
		return &pb.LogoutResponse{Response:pb.ResponseType_FAIL}, errorStatus.Err()
	}
	return &pb.LogoutResponse{Response:pb.ResponseType_SUCCESS}, nil
}

func (s *userApplication) AddFriend(ctx context.Context, request *pb.AddFriendRequest) (*pb.AddFriendResponse, error) {
	if err := s.repository.AddFriend(&model.User{Username:request.UserInfo.Username}, &model.User{Username:request.Friend}); err != nil {
		return &pb.AddFriendResponse{Response:pb.ResponseType_FAIL}, err
	}
	return &pb.AddFriendResponse{Response:pb.ResponseType_SUCCESS}, nil
}

func (s *userApplication) RemoveFriend(ctx context.Context, request *pb.RemoveFriendRequest) (*pb.RemoveFriendResponse, error) {
	if err := s.repository.AddFriend(&model.User{Username:request.UserInfo.Username, Password:request.UserInfo.Username}, &model.User{Username:request.Friend}); err != nil {
		return &pb.RemoveFriendResponse{Response:pb.ResponseType_FAIL}, err
	}
	return &pb.RemoveFriendResponse{Response:pb.ResponseType_SUCCESS}, nil
}


func (s *userApplication) GetFriends(ctx context.Context, request *pb.GetFriendsRequest) (*pb.GetFriendsResponse, error) {
	usr, err := s.repository.GetUserByUsername(request.UserInfo.Username)
	if err != nil {
		errorStatus := status.New(codes.NotFound, fmt.Sprintf("User %s not exists", request.UserInfo.Username))
		return &pb.GetFriendsResponse{}, errorStatus.Err()
	}
	var friendState []*pb.UserState
	for _, friend := range usr.Friends {
		lastTimestamp, _ := ptypes.TimestampProto(friend.Status.LastLoginTime)
		if !friend.Status.Active {
			lastTimestamp, _ = ptypes.TimestampProto(friend.Status.LastLogoutTime)
		}
		friendState = append(friendState, &pb.UserState{Username: friend.Username, Active:friend.Status.Active, Timestamp:lastTimestamp})
	}
	return &pb.GetFriendsResponse{Users: friendState}, nil
}