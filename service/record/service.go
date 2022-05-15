package record

import (
	"KION/domain"
	"context"
	"fmt"
	"time"
)

type RecordService interface {
	CreateRecord(context.Context, domain.Model) error
	GetLatestRecord(context.Context, domain.UserID, domain.VideoID) (time.Duration, error)
}

type RecordStorage interface {
	CreateRecord(context.Context, domain.Model) error
	GetLatestRecord(context.Context, domain.UserID, domain.VideoID) (time.Duration, error)
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
	fmt.Println("Service CreateRecord")
	return nil
}

func (s service) GetLatestRecord(ctx context.Context, userID domain.UserID, videoID domain.VideoID) (time.Duration, error) {
	fmt.Println("Service GetLatestRecord")
	return time.Duration(1), nil
}
