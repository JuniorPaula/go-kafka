package usecases

import "go-message-kafka/internal/entities"

type CreateProductInputDto struct {
	Name  string
	Price float64
}

type CreateProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUsecase struct {
	ProductRepository entities.ProductRepository
}

func (u *CreateProductUsecase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
	product := entities.NewProduct(input.Name, input.Price)
	err := u.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}
	return &CreateProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
