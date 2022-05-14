package record

import (
	"KION/domain"
	"context"
	"time"
)

type RecordService interface {
	CreateRecord(context.Context, domain.Model) error
	GetRecord(context.Context, domain.UserID, domain.VideoID) (time.Duration, error)
}

type RecordStorage interface {
	CreateRecord(context.Context, domain.Model) error
	GetRecord(context.Context, domain.UserID, domain.VideoID) (time.Duration, error)
}

type service struct {
	recordStorage RecordStorage
}

func NewRecordService(recordStorage RecordStorage) RecordService {
	return service{
		recordStorage: recordStorage,
	}
}

func (s service) CreateRecord(ctx context.Context, model domain.Model) error {
	return nil
}

func (s service) GetRecord(ctx context.Context, userID domain.UserID, videoID domain.VideoID) (time.Duration, error) {
	return time.Duration(1), nil
}
