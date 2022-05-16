package api

import (
	"KION/domain"
	record "KION/service/record"
	"KION/specs/gen"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/durationpb"
)

type Controller interface {
	CreateRecord(
		ctx context.Context,
		request *gen.CreateRecordRequest,
	) (*gen.CreateRecordResponse, error)
	GetLatestRecord(
		ctx context.Context,
		request *gen.GetLatestRecordRequest,
	) (*gen.GetLatestRecordResponse, error)
}

type controller struct {
	recordService record.RecordService
}

func NewRecordController(recordService record.RecordService) Controller {
	return &controller{recordService: recordService}
}

func (c controller) CreateRecord(
	ctx context.Context,
	request *gen.CreateRecordRequest,
) (*gen.CreateRecordResponse, error) {
	fmt.Println("Controller CreateRecord")

	videoID, err := convertVideoID(request.GetVideoId())
	if err != nil {
		return nil, err
	}

	userID, err := convertUserID(request.GetUserId())
	if err != nil {
		return nil, err
	}

	videoTime := request.GetTime().AsDuration()
	eventType := request.GetEventType()

	newRecord := domain.NewModel(
		videoID,
		userID,
		videoTime,
		eventType,
	)

	err = c.recordService.CreateRecord(ctx, newRecord)

	var state string
	if err == nil {
		state = "ok"
	} else {
		state = "error during add data to db"
	}

	return &gen.CreateRecordResponse{State: state}, nil
}

func (c controller) GetLatestRecord(
	ctx context.Context,
	request *gen.GetLatestRecordRequest,
) (*gen.GetLatestRecordResponse, error) {
	fmt.Println("Controller GetLatestRecord")

	userID, err := transformUserID(request.GetUserId())
	if err != nil {
		return nil, err
	}

	videoID, err := transformVideoID(request.GetVideoId())
	if err != nil {
		return nil, err
	}

	latestRecord, err := c.recordService.GetLatestRecord(ctx, userID, videoID)
	if err != nil {
		return nil, err
	}

	return &gen.GetLatestRecordResponse{
		Time: durationpb.New(latestRecord),
	}, nil
}
