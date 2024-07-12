package receipt

import (
	"context"
	"github.com/google/uuid"
)

func (i impl) PostPoints(ctx context.Context, id uuid.UUID, points int) error {
	i.idToPoints[id] = points
	return nil
}
