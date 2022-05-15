package clickhouse

import (
	"KION/domain"
	"context"
	"fmt"
	"log"
	"time"

	ch "github.com/ClickHouse/clickhouse-go/v2"
)

// Никита, тебе нужно будет клик как-то тут инициализировать, потому что я хз работает ли клик с этой либой для sql

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

type RecordStorage struct {
	db ch.Conn
}

type Storage interface {
	CreateRecord(context.Context, domain.Model) error
	GetLatestRecord(context.Context, domain.UserID, domain.VideoID) (time.Duration, error)
}

func NewRecordStorage() Storage {
	var (
		ctx     = context.Background()
		db, err = ch.Open(&ch.Options{
			Addr: []string{"localhost:9000"},
			Auth: ch.Auth{
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
	err = db.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Exec(ctx, createTableQuery); err != nil {
		log.Fatal(err)
	}
	return &RecordStorage{db: db}
}

func (r *RecordStorage) CreateRecord(ctx context.Context, model domain.Model) error {
	fmt.Println("Storage CreateRecord")
	return nil
}

func (r *RecordStorage) GetLatestRecord(ctx context.Context, userID domain.UserID, videoID domain.VideoID) (time.Duration, error) {
	fmt.Println("Storage GetLatestRecord")
	return time.Duration(1), nil
}
