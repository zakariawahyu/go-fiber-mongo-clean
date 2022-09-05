package services

import (
	"github.com/zakariawahyu/go-fiber-mongo-clean/entity"
	"github.com/zakariawahyu/go-fiber-mongo-clean/model"
	"github.com/zakariawahyu/go-fiber-mongo-clean/repository"
	validation2 "github.com/zakariawahyu/go-fiber-mongo-clean/validation"
)

type ProductServicesImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductServices(productRepository *repository.ProductRepository) ProductServices {
	return &ProductServicesImpl{
		ProductRepository: *productRepository,
	}
}

func (services *ProductServicesImpl) Create(request model.ProductRequest) (response model.ProductResponse) {
	validation2.ValidateProduct(request)

	product := entity.Product{
		Id:       request.Id,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
	}

	services.ProductRepository.Insert(product)

	response = model.ProductResponse{
		Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
	return response
}
