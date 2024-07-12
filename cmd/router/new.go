package router

import (
	"github.com/hxcuber/receipt-processor/internal/controller/receipt"
	receiptHandler "github.com/hxcuber/receipt-processor/internal/handler/receipt"
)

func New(receiptCtrl receipt.Controller) Router {
	return Router{receiptRESTHandler: receiptHandler.New(receiptCtrl)}
}
