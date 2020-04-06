package main

import (
	"go_grpc_chat/pb"
	"go_grpc_chat/server/chat"
	"go_grpc_chat/server/chat/handler"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":4317")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(handler.LogUnaryServerInterceptor),
						grpc.StreamInterceptor(handler.LogServerStreamInterceptor))
	pb.RegisterChatTaskServer(s, chat.InitChatServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
