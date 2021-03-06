package clickhouse

import (
	"KION/domain"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"

	ch "github.com/ClickHouse/clickhouse-go/v2"
)

const createTableQuery = `
	CREATE TABLE IF NOT EXISTS KION
	(
		VideoId String,
		UserId String,
		EventType String,
		Duration Int,
		EventTime DateTime('Europe/London')
	)

	ENGINE = MergeTree()
	ORDER BY EventTime;
`

const createKafkaTableQuery = `
	CREATE TABLE IF NOT EXISTS KION_queue
	(
		VideoId String,
		UserId String,
		EventType String,
		Duration Int,
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
	SELECT VideoId, UserId, EventType, Duration, EventTime
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

	data := transformModel(model)

	jsonToKafka, err := json.Marshal(data)
	if err != nil {
		return err
	}

	conn, err := kafka.DialLeader(context.Background(), "tcp", "10.244.0.6:9092", "kion-event-topic", 0)
	if err != nil {
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(time.Second * 8))
	conn.WriteMessages(kafka.Message{
		Topic: "kion-event-topic",
		Value: jsonToKafka,
	})

	return nil
}

func (r *RecordStorage) GetLatestRecord(ctx context.Context, userID domain.UserID, videoID domain.VideoID) (time.Duration, error) {
	fmt.Println("Storage GetLatestRecord")

	query := fmt.Sprintf(`
		SELECT Duration FROM KION
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
