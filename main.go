package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/didil/simple-text-rag-go-frontend/server"
	"github.com/didil/simple-text-rag-go-frontend/services"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	grpcaddr := flag.String("grpcaddr", "localhost:50051", "the address to connect to")
	flag.Parse()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	grpcClient, err := services.NewGrpcClient(*grpcaddr)
	if err != nil {
		logger.Fatal("grpc client creation failed", zap.Error(err))
	}
	app := server.NewApp(
		server.WithLogger(logger),
		server.WithGrpcClient(grpcClient),
	)

	r := server.NewRouter(app)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	addr := fmt.Sprintf("%s:%s", host, port)
	httpServer := http.Server{
		Addr:    addr,
		Handler: r,
	}

	logger.Info("listening ...", zap.String("addr", addr))
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal("listener failure", zap.Error(err))
	}
	logger.Info("http server shutdown")
}
