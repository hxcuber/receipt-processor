package receipt

import (
	"context"
	"github.com/google/uuid"
	"github.com/hxcuber/receipt-processor/internal/repository/receipt"
	"github.com/hxcuber/receipt-processor/pkg/model"
)

type Controller interface {
	GetPoints(ctx context.Context, id string) (int, error)
	ProcessReceipts(ctx context.Context, receipt model.Receipt) (uuid.UUID, error)
}

type impl struct {
	repo receipt.Repository
}

func New(r receipt.Repository) Controller {
	return impl{
		repo: r,
	}
}
