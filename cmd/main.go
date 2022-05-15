package main

import (
	"KION/api"
	"KION/service/record"
	"KION/specs/gen"
	"KION/storage/clickhouse"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var db sql.DB

var cont = api.NewRecordController(record.NewRecordService(clickhouse.NewRecordStorage(db)))

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
	conn, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("tcp connection error:", err.Error())
	}

	grpcServer := grpc.NewServer()

	server := goServer{}

	gen.RegisterKionServiceServer(grpcServer, &server)

	fmt.Println("starting server at localhost:8080")
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
