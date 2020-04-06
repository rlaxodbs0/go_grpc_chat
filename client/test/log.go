package test

import (
	"context"

	"google.golang.org/grpc"
	"log"
)

func LogUnaryClientInterceptor(
	ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Println(method)

	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Println(reply)

	return err
}