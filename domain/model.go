package domain

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	videoID   uuid.UUID
	userID    uuid.UUID
	videoTime time.Duration
	eventType Event
}

func NewModel(
	videoID uuid.UUID,
	userID uuid.UUID,
	videoTime time.Duration,
	eventType Event,
) Model {
	return Model{
		videoID:   videoID,
		userID:    userID,
		videoTime: videoTime,
		eventType: eventType,
	}
}

func (m Model) GetVideoID() uuid.UUID {
	return m.videoID
}

func (m Model) GetUserID() uuid.UUID {
	return m.userID
}

func (m Model) GetVideoTime() time.Duration {
	return m.videoTime
}
func (m Model) GetEvent() Event {
	return m.eventType
}
