package application

import (
	"go_grpc_chat/server/chat/domain/repository"
)

type Application struct {
	*AuthApplication
	*ChatApplication
}

func InitApplication(repo repository.ChatRepositoryImpl) *Application{
	return &Application{AuthApplication: InitAuthApplictaion(repo), ChatApplication: InitChatApplication(repo)}
}