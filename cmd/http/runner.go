package main

import (
	"log"
	"net/http"

	"github.com/tiendat-go/common-service/utils/env"
	grpcClient "github.com/tiendat-go/gateway-service/internal/client/grpc"
	grpcCtrl "github.com/tiendat-go/gateway-service/internal/controller/grpc"
)

var (
	httpAddr = env.EnvString("HTTP_ADDR", ":9999")
)

func main() {
	registryClient := grpcClient.NewRegistryClient("localhost:50051", "gateway-service", "9999")
	notificationClient := grpcClient.NewNotificationClient(registryClient)

	notification := grpcCtrl.NewNotificationController(notificationClient)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/notification", notification.GetNotifications)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server!")
	}
}
