package repository

import (
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/domain/model"
	"sync"
)

type ChatRepositoryImpl interface  {
	SignUp(user *model.User) *pb.SignupResponse
	Login(user *model.User) *pb.LoginResponse
	Logout(user *model.User) *pb.LogoutResponse
}

var (
	once sync.Once
	chatRepository *ChatRepositoryImpl
)


func InitChatRepository(cr *ChatRepositoryImpl) {
	once.Do(func() {
		chatRepository = cr
	})
}

