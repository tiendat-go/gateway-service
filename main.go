package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tiendat-go/common-service/utils/env"
)

var (
	httpAddr = env.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	grpcClient, cleanup, err := NewGrpcClient()
	if err != nil {
		log.Fatalf("Failed to connect to gRPC services: %v", err)
	}
	defer cleanup()

	mux := http.NewServeMux()
	handler := NewHandler(grpcClient)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server!")
	}
}
