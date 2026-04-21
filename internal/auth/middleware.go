package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/golang-jwt/jwt/v5"
)

type JWT_middleware struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Skip authentication for login and refresh endpoints
	if info.FullMethod == "/users.UserService/login" || info.FullMethod == "/users.UserService/RefreshToken" {
		return handler(ctx, req)
	}
	// Extract token from metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization header")
	}

	if !strings.HasPrefix(authHeader[0], "Bearer ") {
		return nil, status.Error(codes.Unauthenticated, "authorization header must start with Bearer")
	}

	tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")
	token, err := ValidateToken(tokenStr)
	if err != nil || !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid token claims")
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid user_id in token claims")
	}
	// Add user ID to context
	ctx = context.WithValue(ctx, "user_id", userID)
	return handler(ctx, req)
}
