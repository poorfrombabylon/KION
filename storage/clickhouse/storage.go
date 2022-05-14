package clickhouse

import (
	"KION/domain"
	"context"
	"database/sql"
	"time"
)

// Никита, тебе нужно будет клик как-то тут инициализировать, потому что я хз работает ли клик с этой либой для sql

type RecordStorage struct {
	db sql.DB
}

type Storage interface {
	CreateRecord(context.Context, domain.Model) error
	GetRecord(context.Context, domain.UserID, domain.VideoID) (time.Duration, error)
}

func NewRecordStorage(db sql.DB) Storage {
	return &RecordStorage{db: db}
}

func (r *RecordStorage) CreateRecord(ctx context.Context, model domain.Model) error {
	return nil
}

func (r *RecordStorage) GetRecord(ctx context.Context, userID domain.UserID, videoID domain.VideoID) (time.Duration, error) {
	return time.Duration(1), nil
}
