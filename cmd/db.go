package main

import (
	"context"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

const createTableQuery = `
CREATE TABLE IF NOT EXISTS KION
(
    VideoId String,
    UserId String,
    EventType String,
    EventTime DateTime
)

ENGINE = ReplacingMergeTree()
ORDER BY EventTime
`

func db() {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"localhost:9000"},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: "",
			},
			DialTimeout:     time.Second,
			MaxOpenConns:    10,
			MaxIdleConns:    5,
			ConnMaxLifetime: time.Hour,
		})
	)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := conn.Exec(ctx, createTableQuery); err != nil {
		log.Fatal(err)
	}
}
