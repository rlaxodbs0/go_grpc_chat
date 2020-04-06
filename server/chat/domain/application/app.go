package application

import (
	"go_grpc_chat/server/chat/domain/repository"
)

type Application struct {
	*AuthApplication
	*ChatApplication
}

func InitApplication(userRepo repository.UserRepositoryImpl) *Application{
	return &Application{AuthApplication: InitAuthApplication(userRepo), ChatApplication: InitChatApplication(userRepo)}
}