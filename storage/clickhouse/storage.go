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
		VideoId UUID,
		UserId UUID,
		EventType String,
		EventDuration Int,
		EventTime DateTime('Europe/London')
	)

	ENGINE = MergeTree()
	ORDER BY EventTime;
`

const createKafkaTableQuery = `
	CREATE TABLE IF NOT EXISTS KION_queue
	(
		VideoId UUID,
		UserId UUID,
		EventType String,
		EventDuration Int,
		EventTime DateTime('Europe/London')
	)
	ENGINE = Kafka
	SETTINGS kafka_broker_list = '10.244.0.6:9092',
       		 kafka_topic_list = 'kion-event-topic',
       		 kafka_group_name = 'kion',
       		 kafka_format = 'JSONEachRow';
`

const createMaterializedViewQuery = `
	CREATE MATERIALIZED VIEW IF NOT EXISTS KION_consumer TO KION AS
	SELECT VideoId, UserId, EventType, EventDuration, EventTime
	FROM KION_queue;
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
			Addr: []string{"10.244.0.5:9000"},
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
	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	if err := db.Exec(ctx, createTableQuery); err != nil {
		log.Fatal(err)
	}
	if err := db.Exec(ctx, createKafkaTableQuery); err != nil {
		log.Fatal(err)
	}
	if err := db.Exec(ctx, createMaterializedViewQuery); err != nil {
		log.Fatal(err)
	}
	return &RecordStorage{db: db}
}

func (r *RecordStorage) CreateRecord(ctx context.Context, model domain.Model) error {
	fmt.Println("Storage CreateRecord")

	query := fmt.Sprintf(`
		INSERT INTO KION (VideoId, UserId, EventType, EventDuration, EventTime)
		VALUES ('%s', '%s', '%s', %v, %v);
	`, model.GetVideoID().String(), model.GetUserID().String(), model.GetEvent(), int(model.GetVideoTime().Seconds()), model.GetCreatedAt().Unix())

	err := r.db.Exec(ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (r *RecordStorage) GetLatestRecord(ctx context.Context, userID domain.UserID, videoID domain.VideoID) (time.Duration, error) {
	fmt.Println("Storage GetLatestRecord")

	query := fmt.Sprintf(`
		SELECT EventDuration FROM KION
		WHERE (UserId = '%s' AND VideoId = '%s')
		ORDER BY EventTime DESC
		LIMIT 1
	`, userID.String(), videoID.String())

	var duration int32
	err := r.db.QueryRow(ctx, query).Scan(&duration)

	if err != nil {
		fmt.Println(err.Error())
		return time.Duration(1), err
	}

	return time.Duration(duration), nil
}
