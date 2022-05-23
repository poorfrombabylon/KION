package clickhouse

import (
	"KION/domain"
)

type ModelDTO struct {
	VideoId   string
	UserId    string
	Duration  int
	EventType string
	EventTime string
}

func transformModel(model domain.Model) ModelDTO {

	newModel := ModelDTO{
		VideoId:   model.GetVideoID().String(),
		UserId:    model.GetUserID().String(),
		Duration:  int(model.GetVideoTime().Seconds()),
		EventType: model.GetEvent(),
		EventTime: model.GetCreatedAt().Format("2006-01-02 15:04:05"),
	}

	return newModel
}
