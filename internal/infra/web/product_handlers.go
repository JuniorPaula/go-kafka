package web

import (
	"encoding/json"
	"go-message-kafka/internal/usecases"
	"net/http"
)

type ProductHanlders struct {
	CreateProductUsecase *usecases.CreateProductUsecase
	ListProductsUsecase  *usecases.ListProductsUsecase
}

func NewProductHandlers(createProductUsecase *usecases.CreateProductUsecase, listProductsUsecase *usecases.ListProductsUsecase) *ProductHanlders {
	return &ProductHanlders{
		CreateProductUsecase: createProductUsecase,
		ListProductsUsecase:  listProductsUsecase,
	}
}

func (p *ProductHanlders) CreateProductHanlder(w http.ResponseWriter, r *http.Request) {
	var input usecases.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateProductUsecase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHanlders) ListProductsHanlder(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListProductsUsecase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
