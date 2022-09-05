package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zakariawahyu/go-fiber-mongo-clean/exception"
	"github.com/zakariawahyu/go-fiber-mongo-clean/model"
	"github.com/zakariawahyu/go-fiber-mongo-clean/services"
	"net/http"
)

type ProductController struct {
	ProductServices services.ProductServices
}

func NewProductController(productServices *services.ProductServices) ProductController {
	return ProductController{ProductServices: *productServices}
}

func (controller *ProductController) Routes(app *fiber.App) {
	app.Post("/product", controller.CreateProduct)
}

func (controller *ProductController) CreateProduct(c *fiber.Ctx) error {
	var request model.ProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()
	exception.PanicIfNeeded(err)

	response := controller.ProductServices.Create(request)
	return c.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
