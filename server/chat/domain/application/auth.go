package application

import (
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/domain/model"
	"go_grpc_chat/server/chat/domain/repository"
)

type AuthApplication struct{
	repository repository.UserRepositoryImpl
}

func InitAuthApplication(repo repository.UserRepositoryImpl) *AuthApplication{
	return &AuthApplication{repository: repo}
}

func (s *AuthApplication) SignUp(in *pb.UserInfo) *pb.SignupResponse{
	return s.repository.SignUp(&model.User{Username:in.Username, Password:in.Password})
}

func (s *AuthApplication) Login(in *pb.UserInfo) *pb.LoginResponse{
	return s.repository.Login(&model.User{Username:in.Username, Password:in.Password})
}

func (s *AuthApplication) Logout(in *pb.UserInfo) *pb.LogoutResponse{
	return s.repository.Logout(&model.User{Username:in.Username, Password:in.Password})
}
