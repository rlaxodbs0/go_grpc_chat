package test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

const (
	bucketRate = time.Second
	bucketCap = 5
)

var bucketInstance *bucket
var once sync.Once

type bucket struct {
	capacity uint64
	status   uint64
}

func getBucketInstance() *bucket{
	once.Do(
		func() {
			bucketInstance = &bucket{capacity: bucketCap, status: bucketCap}
			go func() {
				t := time.NewTicker(bucketRate)
				for {
					select {
					case <-t.C:
						log.Println("reset")
						atomic.StoreUint64(&bucketInstance.status, bucketInstance.capacity)
					}
				}
			}()
		})
	return bucketInstance
}

func (b *bucket) Limit() bool {
	for {
		r := atomic.LoadUint64(&b.status)
		log.Println("status : ", r)
		if r == 0 {
			return false
		}
		if !atomic.CompareAndSwapUint64(&b.status, r, r-1) {
			continue
		}
		return true
	}
}



type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	b := getBucketInstance()
	if b.Limit() {
		return w.ClientStream.SendMsg(m)
	}
	return status.Errorf(codes.ResourceExhausted, "Too many request")
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func RateLimitClientStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc,
	cc *grpc.ClientConn, method string,
	streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		s, err := streamer(ctx, desc, cc, method, opts...)
		if err != nil {
			return nil, err
		}
		return newWrappedStream(s), nil

}