package receipt

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/hxcuber/receipt-processor/internal/handler/response/errordesc"
	"github.com/hxcuber/receipt-processor/internal/handler/response/points"
	receiptRepo "github.com/hxcuber/receipt-processor/internal/repository/receipt"
	"github.com/hxcuber/receipt-processor/pkg/util"
	"net/http"
)

func (h Handler) GetPoints() http.HandlerFunc {
	return util.ErrorHandler(
		func(w http.ResponseWriter, r *http.Request) (render.Renderer, int) {
			id := chi.URLParam(r, "id")
			if id == "" {
				return errordesc.ErrorDesc{Description: "id not provided"}, http.StatusBadRequest
			}

			pts, err := h.ctrl.GetPoints(r.Context(), id)
			if err != nil {
				if errors.Is(err, receiptRepo.ErrReceiptNotFound) {
					return errordesc.ErrorDesc{Description: err.Error()}, http.StatusNotFound
				}
				return errordesc.ErrorDesc{Description: err.Error()}, http.StatusInternalServerError
			}

			return points.Points{Points: pts}, http.StatusOK
		})
}
