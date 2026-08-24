package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"rovioletta/todo-list-api/internal/db"
	userSvc "rovioletta/todo-list-api/internal/service/user"
	userGrpc "rovioletta/todo-list-api/internal/transport/grpc/user"
	userPb "rovioletta/todo-list-api/pkg/pb/user"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Starting the Blog app...")

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", slog.String("error", err.Error()))
		return
	}

	dbpool := initDB(logger)
	defer dbpool.Close()

	dbQueries := db.New(db.NewErrorHandler(dbpool))

	// Business logic
	userService := userSvc.NewService(dbQueries)

	conn, err := net.Listen("tcp", ":"+os.Getenv("APP_PORT"))
	if err != nil {
		logger.Error("Failed to run the server", slog.String("error", err.Error()))
		return
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	userPb.RegisterUserServiceServer(grpcServer, userGrpc.NewImplementation(logger, userService))
	reflection.Register(grpcServer)

	grpcServer.Serve(conn)
}

func initDB(logger *slog.Logger) *pgxpool.Pool {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		logger.Error("Database url is not provided")
		os.Exit(1)
	}

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, url)
	if err != nil {
		logger.Error("Unable to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if err := dbpool.Ping(ctx); err != nil {
		logger.Error("Unable to ping database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return dbpool
}
