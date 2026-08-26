package interceptors

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"rovioletta/todo-list-api/internal/pkg/tokens"
)

type ContextKey string

const UserIDKey ContextKey = "user_id"

type AuthInterceptor struct {
	jwtManager *tokens.JWTManager
}

func NewAuthInterceptor(jwtManager *tokens.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager: jwtManager}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {

		if isPublicMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is missing")
		}

		values := md.Get("authorization")
		if len(values) == 0 {
			return nil, status.Error(codes.Unauthenticated, "authorization header is missing")
		}

		authHeader := values[0]
		splits := strings.Split(authHeader, " ")
		if len(splits) != 2 || strings.ToLower(splits[0]) != "bearer" {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization header format")
		}

		accessToken := splits[1]

		claims, err := i.jwtManager.Verify(accessToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "token is invalid: %v", err)
		}

		newCtx := context.WithValue(ctx, UserIDKey, claims.UserID)

		return handler(newCtx, req)
	}
}

func isPublicMethod(fullMethod string) bool {
	publicMethods := map[string]bool{
		"/user.UserService/Login":    true,
		"/user.UserService/Register": true,
	}
	return publicMethods[fullMethod]
}
