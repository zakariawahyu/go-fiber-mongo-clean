package repository

import (
	"github.com/zakariawahyu/go-fiber-mongo-clean/config"
	"github.com/zakariawahyu/go-fiber-mongo-clean/entity"
	"github.com/zakariawahyu/go-fiber-mongo-clean/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewProductRepository(database *mongo.Database) ProductRepository {
	return &ProductRepositoryImpl{
		Collection: database.Collection("products"),
	}
}

func (repository *ProductRepositoryImpl) Insert(product entity.Product) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":      product.Id,
		"name":     product.Name,
		"price":    product.Price,
		"quantity": product.Quantity})
	exception.PanicIfNeeded(err)
}
