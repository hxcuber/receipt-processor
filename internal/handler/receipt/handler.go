package receipt

import (
	"github.com/hxcuber/receipt-processor/internal/controller/receipt"
)

type Handler struct {
	ctrl receipt.Controller
}

func New(c receipt.Controller) Handler {
	return Handler{
		ctrl: c,
	}
}
