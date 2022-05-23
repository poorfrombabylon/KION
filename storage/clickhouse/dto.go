package clickhouse

import (
	"KION/domain"
	"time"
)

type ModelDTO struct {
	VideoID   string
	UserID    string
	VideoTime time.Duration
	EventType string
	CreatedAt time.Time
}

func transformModel(model domain.Model) ModelDTO {

	newModel := ModelDTO{
		VideoID:   model.GetVideoID().String(),
		UserID:    model.GetUserID().String(),
		VideoTime: model.GetVideoTime(),
		EventType: model.GetEvent(),
		CreatedAt: model.GetCreatedAt(),
	}

	return newModel
}
