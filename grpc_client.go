package main

import (
	"log"

	"github.com/tiendat-go/common-service/utils/env"
	pbCore "github.com/tiendat-go/proto-service/gen/core/v1"
	pbCrypto "github.com/tiendat-go/proto-service/gen/crypto/v1"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	Core   pbCore.CoreServiceClient
	Crypto pbCrypto.CryptoServiceClient
}

func NewGrpcClient() (*GrpcClient, func(), error) {
	// Define service addresses
	services := map[string]string{
		"core":   env.EnvString("CORE_SERVICE", "localhost:50051"),
		"crypto": env.EnvString("CRYPTO_SERVICE", "localhost:50052"),
	}

	// Create connections
	connections := make(map[string]*grpc.ClientConn)
	for name, addr := range services {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		connections[name] = conn
		log.Println("Connected to", name, "at", addr)
	}

	// Initialize clients
	clients := &GrpcClient{
		Core:   pbCore.NewCoreServiceClient(connections["core"]),
		Crypto: pbCrypto.NewCryptoServiceClient(connections["crypto"]),
	}

	// Cleanup function to close connections
	cleanup := func() {
		for _, conn := range connections {
			conn.Close()
		}
	}

	return clients, cleanup, nil
}
