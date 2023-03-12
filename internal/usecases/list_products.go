package usecases

import "go-message-kafka/internal/entities"

type ListProductsOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductsUsecase struct {
	ProductRepository entities.ProductRepository
}

func NewListProductsUsecase(productRepository entities.ProductRepository) *ListProductsUsecase {
	return &ListProductsUsecase{ProductRepository: productRepository}
}

func (u *ListProductsUsecase) Execute() ([]*ListProductsOutputDto, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var productsOutput []*ListProductsOutputDto
	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsOutput, nil
}
