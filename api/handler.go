package api

import (
	"context"
	"net/http"
	"phone-verification/data"
	"time"

	"github.com/gofiber/fiber/v2"
)

const appTimeout = time.Second * 10

func (app *Config) sendSMS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		var payload data.OTPData
		if err := c.BodyParser(&payload); err != nil {
			app.errorJSON(c, err)
			return err
		}

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return err
		}

		app.writeJSON(c, fiber.StatusAccepted, "OTP sent successfully")
        return nil
	}
}

func (app *Config) verifySMS() fiber.Handler {
    return func(c *fiber.Ctx) error {
        _, cancel := context.WithTimeout(context.Background(), appTimeout)
        var payload data.VerifyData
        defer cancel()

        app.validateBody(c, &payload)

        newData := data.VerifyData{
            User: payload.User,
            Code: payload.Code,
        }

        err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
        if err != nil {
            app.errorJSON(c, err)
            return err
        }

        app.writeJSON(c, http.StatusAccepted, "OTP verified successfully")
        return nil
    }
}