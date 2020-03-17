package infrastructure

import (
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/domain/model"
	"sync"
	"time"
)

type DB struct{
	UserInfo sync.Map // string - model user
}

type ChatRepository struct {
	db *DB
}

func NewStore() *ChatRepository {
	userInfo := sync.Map{}
	userInfo.Store("a", &model.User{Username:"a", Password:"1"})
	userInfo.Store("b", &model.User{Username:"b", Password:"1"})
	userInfo.Store("c", &model.User{Username:"c", Password:"1"})
	return &ChatRepository{db: &DB{UserInfo: userInfo}}
}

func (ar *ChatRepository) SignUp(user *model.User) *pb.SignupResponse {
	_, ok := ar.db.UserInfo.Load(user.Username); if ok {
		return &pb.SignupResponse{Response:pb.ResponseType_ALREADYEXISTS}
	}
	ar.db.UserInfo.Store(user.Username, &model.User{ Username:user.Username, Password:user.Password})
	return &pb.SignupResponse{Response:pb.ResponseType_SUCCESS}
}

func (ar *ChatRepository) Login(user *model.User) *pb.LoginResponse {
	res, ok := ar.db.UserInfo.Load(user.Username); if !ok {
		return &pb.LoginResponse{Response:pb.ResponseType_NOMATCH}
	}
	if res.(*model.User).Password != user.Password {
		return &pb.LoginResponse{Response:pb.ResponseType_NOMATCH}
	}
	res.(*model.User).Status.Active = true
	res.(*model.User).Status.LoginTime = time.Now()
	return &pb.LoginResponse{Response:pb.ResponseType_SUCCESS}
}

func (ar *ChatRepository) Logout(user *model.User) *pb.LogoutResponse {
	res, ok := ar.db.UserInfo.Load(user.Username); if !ok {
		return &pb.LogoutResponse{Response:pb.ResponseType_NOMATCH}
	}
	res.(*model.User).Status.Active = false
	res.(*model.User).Status.LogoutTime = time.Now()
	return &pb.LogoutResponse{Response:pb.ResponseType_SUCCESS}
}

func (ar *ChatRepository) GetUserByUsername(username string) *model.User{
	res, ok := ar.db.UserInfo.Load(username); if ok {
		return res.(*model.User)
	}
	return nil
}

func (ar *ChatRepository) GetActiveUserPointerSlice(username string) []*model.User{
	activeUserPointerSlice := []*model.User{}
	ar.db.UserInfo.Range(func(k, user interface{}) bool {
		if user.(*model.User).Status.Active {
			activeUserPointerSlice = append(activeUserPointerSlice, user.(*model.User))
		}
		return true
	})
	return activeUserPointerSlice
}

