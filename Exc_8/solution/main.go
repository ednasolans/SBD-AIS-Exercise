package main

import (
	"exc8/client"
	"exc8/server"
	"log"
	"time"
)

func main() {
	// Start the gRPC server in a separate goroutnine
	go func() {
		// todo start server
		if err := server.StartGrpcServer(); err != nil {
			log.Fatalf("failed to start server!!: %v", err)
		}

	}()

	time.Sleep(1 * time.Second)

	// todo start client
	cli, err := client.NewGrpcClient()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	if err := cli.Run(); err != nil {
		log.Fatalf("client error: %v", err)
	}

	println("Orders complete!")
}
