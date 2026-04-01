package orders

import (
	"log"
	"net/http"

	"github.com/sorrawit2546/internal/json"
)

type OrderHandler struct {
	service IOrderService
}

func NewOrderHandler(s IOrderService) *OrderHandler {
	return &OrderHandler{
		service: s,
	}
}

func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	data := createOrderParams{}
	if err := json.Read(r, &data); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.service.placeOrder(r.Context(), data)
	json.Write(w, http.StatusCreated, nil)
}
