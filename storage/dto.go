package storage

import (
	"github.com/google/uuid"
	"time"
)

type RecordDb struct {
	video_id   uuid.UUID `db:"video_id"`
	user_id    uuid.UUID `db:"user_id"`
	event_type string    `db:"event_type"`
	time       time.Time `db:"time"`
}
