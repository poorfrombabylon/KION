package main

import (
	"context"
	"log"
	"time"

	"net/http"

	"github.com/ClickHouse/clickhouse-go/v2"
)

var UrlHello = []byte(`
<html>
	<body>
		<form action = "/" method = "POST">
			Enter Your Name: <input type="text" name="userName">
			<input type="submit" value="ENTER">
		</form>
	</body>
	<body>
		
	</body>
`)

const createTableQuery = `
	CREATE TABLE IF NOT EXISTS KION
	(
		VideoId String,
		UserId String,
		EventType String,
		EventTime DateTime
	)

	ENGINE = ReplacingMergeTree()
	ORDER BY EventTime;
`

func handle228(w http.ResponseWriter, r *http.Request) {
	w.Write(UrlHello)
	name := r.FormValue("userName")
	if name != "" {
		w.Write([]byte("Hello " + name))
	}
}

func main() {
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

	http.HandleFunc("/", handle228)

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})

	http.ListenAndServe(":8080", nil)
}
