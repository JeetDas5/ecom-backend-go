package products

import (
	"github/JeetDas5/ecom-app/internal/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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

func (h *handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByID(r.Context(), id)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "product not found", http.StatusNotFound)
			return
		}

		log.Printf("error getting product by id %d: %s", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
