package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"time"
)

func logInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)
	grpclog.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)
	return h, err
}

func WithLogInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(logInterceptor)
}
