package interceptors

import (
	"context"
	"log/slog"

	"buf.build/go/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ValidationInterceptor struct {
	logger *slog.Logger
	v      protovalidate.Validator
}

func NewValidationInterceptor(logger *slog.Logger) *ValidationInterceptor {
	v, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	return &ValidationInterceptor{
		v:      v,
		logger: logger,
	}
}

func (i *ValidationInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		
		if msg, ok := req.(proto.Message); ok {
			if err := i.v.Validate(msg); err != nil {
				i.logger.Debug("i.v.Validate", slog.String("error", err.Error()))
				return nil, status.Errorf(codes.InvalidArgument, "validation error")
			}
		}

		return handler(ctx, req)
	}
}
