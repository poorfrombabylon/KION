package clickhouse

import (
	"KION/domain"
	"context"
	"fmt"
	"log"
	"time"

	ch "github.com/ClickHouse/clickhouse-go/v2"
)

const createTableQuery = `
	CREATE TABLE IF NOT EXISTS KION
	(
		VideoId String,
		UserId String,
		EventType String,
		EventTime Int
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
			Addr: []string{"212.23.220.55:9000"},
			Auth: ch.Auth{
				Database: "default",
				Username: "clickhouse_operator",
				Password: "clickhouse_operator_password",
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

	query := fmt.Sprintf(`
		INSERT INTO KION (VideoId, UserId, EventType, EventTime)
		VALUES ('%s', '%s', '%s', %v);
	`, model.GetVideoID().String(), model.GetUserID().String(), model.GetEvent(), int(model.GetVideoTime().Seconds()))

	err := r.db.Exec(ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (r *RecordStorage) GetLatestRecord(ctx context.Context, userID domain.UserID, videoID domain.VideoID) (time.Duration, error) {
	fmt.Println("Storage GetLatestRecord")
	return time.Duration(1), nil
}
