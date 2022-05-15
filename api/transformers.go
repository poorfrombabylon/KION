package api

import (
	"KION/domain"
	"github.com/google/uuid"
	"github.com/ignishub/terr"
)

var (
	errParseVideoIDToUuid = terr.BadRequest(
		"INVALID_VIDEO_UUID",
		"некорректный формат UUID видеозаписи",
	)

	errParseUserIDToUuid = terr.BadRequest(
		"INVALID_USER_UUID",
		"некорректный формат UUID пользователя",
	)
)

func convertVideoID(id string) (domain.VideoID, error) {
	if id == "" {
		return domain.VideoID{}, errParseVideoIDToUuid
	}

	videoUUID, err := uuid.Parse(id)
	if err != nil {
		return domain.VideoID{}, errParseVideoIDToUuid
	}

	videoID := domain.VideoID(videoUUID)

	return videoID, nil
}

func convertUserID(id string) (domain.UserID, error) {
	if id == "" {
		return domain.UserID{}, errParseUserIDToUuid
	}

	userUUID, err := uuid.Parse(id)
	if err != nil {
		return domain.UserID{}, errParseUserIDToUuid
	}

	userID := domain.UserID(userUUID)

	return userID, nil
}

func convertEvent(event string) domain.Event {
	switch event {
	case "stop":
		return domain.StopVideoEvent
	case "reverse":
		return domain.RevertVideoEvent
	case "forward":
		return domain.ForwardVideoEvent
	case "exit":
		return domain.ExitVideoEvent
	default:
		return domain.Nothing
	}
}

func transformVideoID(id string) (domain.VideoID, error) {
	if id == "" {
		return domain.VideoID{}, errParseVideoIDToUuid
	}

	videoUUID, err := uuid.Parse(id)
	if err != nil {
		return domain.VideoID{}, errParseVideoIDToUuid
	}

	videoID := domain.VideoID(videoUUID)

	return videoID, nil
}

func transformUserID(id string) (domain.UserID, error) {
	if id == "" {
		return domain.UserID{}, errParseUserIDToUuid
	}

	userUUID, err := uuid.Parse(id)
	if err != nil {
		return domain.UserID{}, errParseUserIDToUuid
	}

	userID := domain.UserID(userUUID)

	return userID, nil
}
