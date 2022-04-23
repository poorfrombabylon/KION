package domain

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	videoID   uuid.UUID
	userID    uuid.UUID
	videoTime string
	eventType Event
	eventTime time.Time
}

func (m Model) GetVideoID() uuid.UUID {
	return m.videoID
}

func (m Model) GetUserID() uuid.UUID {
	return m.userID
}

func (m Model) GetVideoTime() string {
	return m.videoTime
}
func (m Model) GetEvent() Event {
	return m.eventType
}

func (m Model) GetEventTime() time.Time {
	return m.eventTime
}
