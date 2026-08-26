package main

import (
	"context"
	"log/slog"
	"net"
	"os"
	"time"

	"rovioletta/todo-list-api/internal/db"
	"rovioletta/todo-list-api/internal/pkg/tokens"
	taskSvc "rovioletta/todo-list-api/internal/service/task"
	userSvc "rovioletta/todo-list-api/internal/service/user"
	"rovioletta/todo-list-api/internal/transport/grpc/interceptors"
	taskGrpc "rovioletta/todo-list-api/internal/transport/grpc/task"
	userGrpc "rovioletta/todo-list-api/internal/transport/grpc/user"
	taskPb "rovioletta/todo-list-api/pkg/pb/task"
	userPb "rovioletta/todo-list-api/pkg/pb/user"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Define logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Info("Starting the Blog app...")

	// Read environment variables
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", slog.String("error", err.Error()))
		return
	}

	// Postgres queries
	dbpool := initDB(logger)
	defer dbpool.Close()

	dbQueries := db.NewErrorHandler(db.New(dbpool))

	// Tokens logic
	tokensManager := tokens.NewJWTManager(os.Getenv("TOKENS_SECRET_KEY"), time.Duration(time.Hour*12))

	// Business logic service
	userService := userSvc.NewService(dbQueries, tokensManager)
	taskService := taskSvc.NewService(dbQueries)

	// Run server
	conn, err := net.Listen("tcp", ":"+os.Getenv("APP_PORT"))
	if err != nil {
		logger.Error("Failed to run the server", slog.String("error", err.Error()))
		return
	}

	authInterceptor := interceptors.NewAuthInterceptor(tokensManager)
	validationInterceptor := interceptors.NewValidationInterceptor(logger)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			authInterceptor.Unary(),
			validationInterceptor.Unary(),
		),
	)

	userPb.RegisterUserServiceServer(grpcServer, userGrpc.NewImplementation(logger, userService))
	taskPb.RegisterTaskServiceServer(grpcServer, taskGrpc.NewImplementation(logger, taskService))
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
