package application

import (
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/domain/model"
	"go_grpc_chat/server/chat/domain/repository"
	"log"
)

type AuthApplictaion struct{
	repository repository.ChatRepositoryImpl
}

func InitAuthApplictaion(repo repository.ChatRepositoryImpl) *AuthApplictaion{
	log.Print("init auth app")
	return &AuthApplictaion{repository: repo}
}

func (s *AuthApplictaion) SignUp(in *pb.UserInfo) *pb.SignupResponse{
	return s.repository.SignUp(&model.User{Username:in.Username, Password:in.Password})
}

func (s *AuthApplictaion) Login(in *pb.UserInfo) *pb.LoginResponse{
	return s.repository.Login(&model.User{Username:in.Username, Password:in.Password})
}

func (s *AuthApplictaion) Logout(in *pb.UserInfo) *pb.LogoutResponse{
	return s.repository.Logout(&model.User{Username:in.Username, Password:in.Password})
}
