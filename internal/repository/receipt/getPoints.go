package receipt

import (
	"context"
	"github.com/google/uuid"
)

func (i impl) GetPoints(ctx context.Context, id uuid.UUID) (int, error) {
	v, ok := i.idToPoints[id]
	if !ok {
		return 0, ErrReceiptNotFound
	}
	return v, nil
}
