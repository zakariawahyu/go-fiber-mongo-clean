package repository

import "github.com/zakariawahyu/go-fiber-mongo-clean/entity"

type ProductRepository interface {
	Insert(product entity.Product)
	SelectAll() (product []entity.Product)
	SelectById(productId string) entity.Product
}
