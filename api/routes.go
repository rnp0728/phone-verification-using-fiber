package api

import (
	"github.com/gofiber/fiber/v2"
)
type Config struct {
	Router *fiber.App
}

func (app *Config) Routes() {
    app.Router.Post("/otp", app.sendSMS())
    app.Router.Post("/verifyOTP", app.verifySMS())
}