package receipt

import (
	"github.com/go-chi/render"
	"log"
	"net/http"
	"github.com/hxcuber/receipt-processor/internal/handler/request/receipt"
	"github.com/hxcuber/receipt-processor/internal/handler/response/errordesc"
	"github.com/hxcuber/receipt-processor/internal/handler/response/id"
	"github.com/hxcuber/receipt-processor/pkg/model"
	"github.com/hxcuber/receipt-processor/pkg/util"
	"time"
)

func (h Handler) ProcessReceipts() http.HandlerFunc {
	return util.ErrorHandler(func(w http.ResponseWriter, r *http.Request) (render.Renderer, int) {
		rec := receipt.Receipt{}
		err := render.Bind(r, &rec)
		if err != nil {
			return errordesc.ErrorDesc{Description: "unable to bind request"}, http.StatusBadRequest
		}
		log.Println(rec)
		modelDate, err := time.Parse(time.DateOnly, rec.PurchaseDate)
		if err != nil {
			return errordesc.ErrorDesc{Description: "unable to parse receipt date"}, http.StatusInternalServerError
		}
		modelTime, err := time.Parse("15:04", rec.PurchaseTime)
		if err != nil {
			return errordesc.ErrorDesc{Description: "unable to parse receipt time"}, http.StatusInternalServerError
		}

		var modelItems []model.Item
		for _, v := range rec.Items {
			modelItems = append(modelItems, model.Item{
				ShortDescription: v.ShortDescription,
				Price:            float64(v.Price),
			})
		}

		modelRec := model.Receipt{
			Retailer:     rec.Retailer,
			PurchaseDate: modelDate,
			PurchaseTime: modelTime,
			Items:        modelItems,
			Total:        float64(rec.Total),
		}

		newId, err := h.ctrl.ProcessReceipts(r.Context(), modelRec)
		if err != nil {
			return errordesc.ErrorDesc{Description: "something went wrong"}, http.StatusInternalServerError
		}

		return id.Id{Id: newId.String()}, http.StatusOK
	})
}
