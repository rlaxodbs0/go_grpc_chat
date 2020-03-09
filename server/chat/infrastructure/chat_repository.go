package infrastructure

import (
	"go_grpc_chat/server/chat/domain/model"
	"sync"
	"go_grpc_chat/pb"
)

type DB struct{
	UserInfo sync.Map // string - model user
}

type ChatRepository struct {
	db *DB
}

func NewStore() *ChatRepository {
	return &ChatRepository{db: &DB{UserInfo: sync.Map{}}}
}

func (ar *ChatRepository) SignUp(user *model.User) *pb.SignupResponse {
	_, ok := ar.db.UserInfo.Load(user.Username); if ok {
		return &pb.SignupResponse{Response:pb.ResponseType_ALREADYEXISTS}
	}
	ar.db.UserInfo.Store(user.Username, model.User{ Username:user.Username, Password:user.Password})
	return &pb.SignupResponse{Response:pb.ResponseType_SUCCESS}
}

func (ar *ChatRepository) Login(user *model.User) *pb.LoginResponse {
	_, ok := ar.db.UserInfo.Load(user.Username); if !ok {
		return &pb.LoginResponse{Response:pb.ResponseType_NOMATCH}
	}
	return &pb.LoginResponse{Response:pb.ResponseType_SUCCESS}
}

func (ar *ChatRepository) Logout(user *model.User) *pb.LogoutResponse {
	_, ok := ar.db.UserInfo.Load(user.Username); if ok {
		return &pb.LogoutResponse{Response:pb.ResponseType_NOMATCH}
	}
	return &pb.LogoutResponse{Response:pb.ResponseType_SUCCESS}
}

