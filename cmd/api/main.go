package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers/person"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/config"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	service := solid.NewSolidService(c.SolidEnv, c.SolidAPIKey)

	person.RegisterRoutes(app, service)

	err = app.Listen(c.Port)
	if err != nil {
		log.Fatalln("Failed at Listen", err)
	}
}
