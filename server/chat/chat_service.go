package chat

import (
	"go_grpc_chat/server/chat/domain/application"
	"go_grpc_chat/server/chat/handler"
	"go_grpc_chat/server/chat/infrastructure"
)

func InitChatServer() {
	port := ":4317"
	handler.InitGrpcHandler(port,
		application.InitApplictaion(infrastructure.NewStore()))
}
