package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserID uuid.UUID

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

type VideoID uuid.UUID

func (v VideoID) String() string {
	return uuid.UUID(v).String()
}

type Model struct {
	videoID   VideoID
	userID    UserID
	videoTime time.Duration
	eventType string
	createdAt time.Time
}

func NewModel(
	videoID VideoID,
	userID UserID,
	videoTime time.Duration,
	eventType string,
) Model {
	return Model{
		videoID:   videoID,
		userID:    userID,
		videoTime: videoTime,
		eventType: eventType,
		createdAt: time.Now().In(time.UTC),
	}
}

func (m Model) GetVideoID() VideoID {
	return m.videoID
}

func (m Model) GetUserID() UserID {
	return m.userID
}

func (m Model) GetVideoTime() time.Duration {
	return m.videoTime
}
func (m Model) GetEvent() string {
	return m.eventType
}

func (m Model) GetCreatedAt() time.Time {
	return m.createdAt
}
