package infrastructure

import (
	"errors"
	"fmt"
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/domain/model"
	"go_grpc_chat/server/chat/domain/repository"
	"sync"
	"time"
)

type DB struct{
	info sync.Map // string - model user
}

type UserRepository struct {
	userMap *DB
}

var _ repository.UserRepositoryImpl = &UserRepository{} //interface impl

func NewUserRepository() *UserRepository {
	info := sync.Map{}
	info.Store("g1", &model.User{Username:"g1", Password:"g1"})
	info.Store("g2", &model.User{Username:"g2", Password:"g2"})
	return &UserRepository{userMap: &DB{info: info}}
}

func (ar *UserRepository) SignUp(user *model.User) *pb.SignupResponse {
	_, ok := ar.userMap.info.Load(user.Username); if ok {
		return &pb.SignupResponse{Response:pb.ResponseType_ALREADYEXISTS}
	}
	ar.userMap.info.Store(user.Username, &model.User{ Username:user.Username, Password:user.Password})
	return &pb.SignupResponse{Response:pb.ResponseType_SUCCESS}
}

func (ar *UserRepository) Login(user *model.User) *pb.LoginResponse {
	res, ok := ar.userMap.info.Load(user.Username); if !ok {
		return &pb.LoginResponse{Response:pb.ResponseType_NOMATCH}
	}
	if res.(*model.User).Password != user.Password {
		return &pb.LoginResponse{Response:pb.ResponseType_NOMATCH}
	}
	res.(*model.User).Status.Active = true
	res.(*model.User).Status.LoginTime = time.Now()
	return &pb.LoginResponse{Response:pb.ResponseType_SUCCESS}
}

func (ar *UserRepository) Logout(user *model.User) *pb.LogoutResponse {
	res, ok := ar.userMap.info.Load(user.Username); if !ok {
		return &pb.LogoutResponse{Response:pb.ResponseType_NOMATCH}
	}
	res.(*model.User).Status.Active = false
	res.(*model.User).Status.LogoutTime = time.Now()
	return &pb.LogoutResponse{Response:pb.ResponseType_SUCCESS}
}

func (ar *UserRepository) GetUserByUsername(username string) *model.User{
	res, ok := ar.userMap.info.Load(username); if ok {
		return res.(*model.User)
	}
	return nil
}

func (ar *UserRepository) GetActiveUserPointerSlice(username string) []*model.User{
	var activeUserPointerSlice []*model.User
	ar.userMap.info.Range(func(k, user interface{}) bool {
		if user.(*model.User).Status.Active {
			activeUserPointerSlice = append(activeUserPointerSlice, user.(*model.User))
		}
		return true
	})
	return activeUserPointerSlice
}

func (ar *UserRepository) UpdateUser(user *model.User) error {
	_, ok := ar.userMap.info.Load(user.Username); if ok {
		ar.userMap.info.Delete(user.Username)
		ar.userMap.info.Store(user.Username, user)
		return nil
	}
	return errors.New(fmt.Sprintf("No User %v", user.Username))
}

