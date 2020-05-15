package repository

import (
	"go_grpc_chat/server/chat/domain/model"
)

type UserRepositoryImpl interface  {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	AddFriend(user *model.User, friend *model.User) error
	RemoveFriend(user *model.User, friend *model.User) error
}
