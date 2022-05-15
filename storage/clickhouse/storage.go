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
		RecordId Int NOT NULL AUTO_INCREMENT,
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

	query := fmt.Sprintf(`
		INSERT INTO KION (VideoId, UserId, EventType, EventTime)
		VALUES (%s, %s, %s, %s);
	`, model.GetVideoID().String(), model.GetUserID().String(), model.GetEvent().String(), model.GetVideoTime())

	err := r.db.Exec(ctx, query)
	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}

	return nil
}

func (r *RecordStorage) GetLatestRecord(ctx context.Context, userID domain.UserID, videoID domain.VideoID) (time.Duration, error) {
	fmt.Println("Storage GetLatestRecord")
	return time.Duration(1), nil
}
