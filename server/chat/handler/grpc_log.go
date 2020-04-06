package handler

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func LogUnaryServerInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println(info.FullMethod)
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("error %s occurred!", err)
	}
	return m, err
}


type wrappedStream struct {
	grpc.ServerStream
}


func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("[RECV] %v", m)
	return w.ServerStream.RecvMsg(m)
}


func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("[SEND] %v", m)
	return w.ServerStream.SendMsg(m)
}


func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}


func LogServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println(info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("error %s occurred!", err)
	}
	return err
}
