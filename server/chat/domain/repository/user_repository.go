package repository

import (
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/domain/model"
)

type UserRepositoryImpl interface  {
	GetUserByUsername(username string) *model.User
	SignUp(user *model.User) *pb.SignupResponse
	Login(user *model.User) *pb.LoginResponse
	Logout(user *model.User) *pb.LogoutResponse
	GetActiveUserPointerSlice(username string) []*model.User
}
