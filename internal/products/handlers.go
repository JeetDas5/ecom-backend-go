package products

import (
	"github/JeetDas5/ecom-app/internal/json"
	"log"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Printf("error listing products: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// products := struct {
	// 	Products []string `json:"products"`
	// }{}
	json.Write(w, http.StatusOK, products)
}
