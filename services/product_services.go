package services

import "github.com/zakariawahyu/go-fiber-mongo-clean/model"

type ProductServices interface {
	Create(request model.ProductRequest) (response model.ProductResponse)
}
