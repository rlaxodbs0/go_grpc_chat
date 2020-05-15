package chat

import (
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat/handler"
	"google.golang.org/grpc"
	"log"
	"net"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

func StartChatServer() {
	lis, err := net.Listen("tcp", ":4317")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		handler.LogUnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(),
	)), grpc.StreamInterceptor(handler.LogServerStreamInterceptor))
	pb.RegisterChatServiceServer(s, &handler.ChatServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}