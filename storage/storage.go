package storage

import (
	"context"
)

type IStorage interface {
	CreateRecord(ctx context.Context)
}
