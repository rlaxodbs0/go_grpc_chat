package application

import (
	"go_grpc_chat/server/chat/domain/repository"
)

type Application struct {
	*AuthApplictaion
	*ChatApplictaion
}

func InitApplictaion(repo repository.ChatRepositoryImpl) *Application{
	return &Application{AuthApplictaion: InitAuthApplictaion(repo), ChatApplictaion: InitChatApplictaion(repo)}
}