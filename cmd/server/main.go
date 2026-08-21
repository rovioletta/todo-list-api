package main

import (
	"log/slog"
	"net"
	"os"

	"rovioletta/todo-list-api/pkg/pb/user"
	user_srv "rovioletta/todo-list-api/internal/app/user"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("Starting the Blog app...")

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", slog.String("error", err.Error()))
		return
	}
	
	conn, err := net.Listen("tcp", os.Getenv("APP_ADDRESS"))
	if err != nil {
		logger.Error("Failed to run the server", slog.String("error", err.Error()))
		return
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	user.RegisterAuthServiceServer(grpcServer, user_srv.NewUserService())
	grpcServer.Serve(conn)
}
