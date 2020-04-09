package chat

import (
	"go_grpc_chat/server/chat/domain/application"
	"go_grpc_chat/server/chat/handler"
	"go_grpc_chat/server/chat/infrastructure"
)

func InitChatServer() *handler.ChatServer {
	return handler.InitGrpcHandler(application.InitApplication(infrastructure.NewUserRepository()))
}
