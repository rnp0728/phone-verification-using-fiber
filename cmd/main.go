package main

import (
	"phone-verification/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
    router := fiber.New()

    //initialize config
    app := api.Config{Router: router}

    //routes
    app.Routes()

    router.Listen(":8080")
}