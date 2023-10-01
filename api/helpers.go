package api

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


type jsonResponse struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
    Data    any    `json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(c *fiber.Ctx, data any) error {
    //validate the request body
    if err := c.BodyParser(&data); err != nil {
        return err
    }
    //use the validator library to validate required fields
    if err := validate.Struct(&data); err != nil {
        return err
    }
    return nil
}

func (app *Config) writeJSON(c *fiber.Ctx, status int, data any) {
    c.Status(status).JSON(jsonResponse{Status: status, Message: "success", Data: data})
}

func (app *Config) errorJSON(c *fiber.Ctx, err error, status ...int) {
    statusCode := http.StatusBadRequest
    if len(status) > 0 {
        statusCode = status[0]
    }
    c.Status(statusCode).JSON(jsonResponse{Status: statusCode, Message: err.Error()})
}