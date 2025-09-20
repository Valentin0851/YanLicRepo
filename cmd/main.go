package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"other_files/commonlibsLesson/internal/config"
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

	cfg, err := config.New()
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to load config", zap.Error(err))
	}

	_, err = postgres.New(cfg.Postgres)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Error(ctx, "failed to connect to db", zap.Error(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.GRPCPort))
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
