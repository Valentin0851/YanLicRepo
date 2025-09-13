package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"other_files/commonlibsLesson/internal/service"
	test "other_files/commonlibsLesson/pkg/api/test/api"
	"other_files/commonlibsLesson/pkg/logger"
	"other_files/commonlibsLesson/pkg/postgres"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)

	pgCfg := postgres.Config{
		Host:     "localhost",
		Port:     "5431",
		Username: "root",
		Password: "1234",
		Database: "postgres",
	}

	db, err := postgres.New(pgCfg)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Error(ctx, "failed to connect to db", zap.Error(err))
	}

	fmt.Println(db.Ping(ctx))

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := service.New()
	server := grpc.NewServer()
	test.RegisterOrderServiceServer(server, srv)
	if err := server.Serve(lis); err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve: ", zap.Error(err))
	}
}
