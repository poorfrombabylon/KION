package api

import (
	"KION/domain"
	record "KION/service/record"
	"KION/specs/gen"
	"context"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Controller interface {
	CreateRecord(
		ctx context.Context,
		request *gen.CreateRecordRequest,
	) (*gen.CreateRecordResponse, error)
	GetRecord(
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
	videoID, err := convertVideoID(request.GetVideoId())
	if err != nil {
		return nil, err
	}

	userID, err := convertUserID(request.GetUserId())
	if err != nil {
		return nil, err
	}

	videoTime := request.GetTime().AsDuration()
	eventType := convertEvent(request.GetEventType())

	newRecord := domain.NewModel(
		videoID,
		userID,
		videoTime,
		eventType,
	)

	err = c.recordService.CreateRecord(ctx, newRecord)

	var state string
	switch err {
	case nil:
		state = "ok"
	default:
		state = "error during add data to db"
	}
	return &gen.CreateRecordResponse{State: state}, nil
}

func (c controller) GetRecord(
	ctx context.Context,
	request *gen.GetLatestRecordRequest,
) (*gen.GetLatestRecordResponse, error) {
	userID, err := transformUserID(request.GetUserId())
	if err != nil {
		return nil, err
	}

	videoID, err := transformVideoID(request.GetVideoId())
	if err != nil {
		return nil, err
	}

	latestRecord, err := c.recordService.GetRecord(ctx, userID, videoID)
	if err != nil {
		return nil, err
	}

	return &gen.GetLatestRecordResponse{
		Time: durationpb.New(latestRecord),
	}, nil
}
