package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "لا يوجد بيانات Metadata")
	}

	//fmt.Printf("Incoming Metadata: %v\n", md) // Debugging line to print incoming metadata

	values := md["authorization"]
	if len(values) == 0 || values[0] != "secret-key" {
		return nil, status.Errorf(codes.Unauthenticated, "التوكن غير صحيح أو مفقود")
	}

	return handler(ctx, req)
}
