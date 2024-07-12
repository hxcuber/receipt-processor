package receipt

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	GetPoints(ctx context.Context, id uuid.UUID) (int, error)
	PostPoints(ctx context.Context, id uuid.UUID, points int) error
}

type impl struct {
	idToPoints map[uuid.UUID]int
}

func New() Repository {
	return impl{
		idToPoints: make(map[uuid.UUID]int),
	}
}
