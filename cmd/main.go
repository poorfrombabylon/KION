package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx := context.Background()
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	} else {
		fmt.Println("погнали")
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		<-ctx.Done()
		server.GracefulStop()

		return nil
	})
}
