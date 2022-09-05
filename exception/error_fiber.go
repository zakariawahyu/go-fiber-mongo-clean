package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-mongo-clean/model"
	"net/http"
)

func ErrorHandler(ctx fiber.Ctx, err error) error {
	_, ok := err.(ValidationErr)
	if ok {
		return ctx.JSON(model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
