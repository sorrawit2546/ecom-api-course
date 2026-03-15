package products

import (
	"log"
	"net/http"

	"github.com/sorrawit2546/internal/json"
)

type ProductsInterface interface {
}

type Products struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type handler struct {
	service IService
}

func NewHandler(service IService) *handler {
	return &handler{
		service: service,
	}
}

// create ListProduct
func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	//call service
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	json.Write(w, http.StatusOK, products)
}
