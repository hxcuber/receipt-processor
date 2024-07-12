package receipt

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/receipt-processor/pkg/logerr"
	"github.com/hxcuber/receipt-processor/pkg/model"
	"github.com/hxcuber/receipt-processor/pkg/points"
)

func (i impl) ProcessReceipts(ctx context.Context, receipt model.Receipt) (uuid.UUID, error) {
	const funcName = "ProcessReceipts"
	id := uuid.New()
	pts := points.CalcPoints(receipt)

	err := i.repo.PostPoints(context.Background(), id, pts)
	if err != nil {
		logerr.LogErrMessage(funcName, "creating new receipt entry for id %s", err, id.String())
		return uuid.Nil, err
	}

	return id, nil
}
