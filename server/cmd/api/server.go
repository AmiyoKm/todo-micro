package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/AmiyoKm/todo-micro/configs"
	"github.com/AmiyoKm/todo-micro/internal/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func server(cfg configs.Config, svc handler.Service) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	go func() {
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		log.Println("Received shutdown signal...")
		cancel()
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	handler.NewGRPCHandler(grpcServer, svc)

	reflection.Register(grpcServer)

	log.Printf("Starting gRPC server Todo service on port %s", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	// wait for the shutdown signal
	<-ctx.Done()
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()
}
