package main

import (
	"KION/api"
	"KION/service/record"
	"KION/specs/gen"
	"KION/storage/clickhouse"
	"context"
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

var db sql.DB

var cont = api.NewRecordController(record.NewRecordService(clickhouse.NewRecordStorage()))

type goServer struct {
	gen.UnimplementedKionServiceServer
}

func (s *goServer) CreateRecord(ctx context.Context, req *gen.CreateRecordRequest) (*gen.CreateRecordResponse, error) {
	return cont.CreateRecord(ctx, req)
}

func (s *goServer) GetLatestRecord(ctx context.Context, req *gen.GetLatestRecordRequest) (*gen.GetLatestRecordResponse, error) {
	return cont.GetLatestRecord(ctx, req)
}

func main() {
	ctx := context.Background()

	fmt.Println("starting server at localhost:8888")
	go initHttpServer()

	fmt.Println("starting server at localhost:8082")
	initGrpcServer(ctx)

}

func initGrpcServer(ctx context.Context) error {
	conn, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatal("tcp connection error:", err.Error())
	}

	grpcServer := grpc.NewServer()

	server := goServer{}

	gen.RegisterKionServiceServer(grpcServer, &server)

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := grpcServer.Serve(conn); err != nil {
			return fmt.Errorf("failed to serve gRPC server: %w", err)
		}

		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		grpcServer.GracefulStop()

		return nil
	})

	return group.Wait()
}

func initHttpServer() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8888", nil)
}
