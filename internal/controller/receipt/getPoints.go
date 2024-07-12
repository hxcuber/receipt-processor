package receipt

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/receipt-processor/pkg/logerr"
	"log"
)

func (i impl) GetPoints(ctx context.Context, id string) (int, error) {
	const funcName = "GetPoints"
	parsed, err := uuid.Parse(id)
	if err != nil {
		log.Print(logerr.LogErrMessage(funcName, "parsing id", err))
		return 0, err
	}

	points, err := i.repo.GetPoints(context.Background(), parsed)
	if err != nil {
		log.Print(logerr.LogErrMessage(funcName, "retrieving points from id %s", err, parsed.String()))
		return 0, err
	}

	return points, nil
}
