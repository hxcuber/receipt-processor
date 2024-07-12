package main

import (
	"log"
	"net/http"
	"github.com/hxcuber/receipt-processor/cmd/router"
	receiptCtrl "github.com/hxcuber/receipt-processor/internal/controller/receipt"
	receiptRepo "github.com/hxcuber/receipt-processor/internal/repository/receipt"
	"time"
)

func main() {
	recRepo := receiptRepo.New()
	recCtrl := receiptCtrl.New(recRepo)
	handler := router.New(recCtrl).Handler()

	srv := http.Server{
		Addr:         ":3000",
		Handler:      handler,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	log.Printf("Server starting up on http://localhost%s", srv.Addr)
	srv.ListenAndServe()
}
