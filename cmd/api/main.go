package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	bankingController "github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers/banking"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers/person"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/config"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
	bankingService "github.com/luke92/solid-fintech-integration-golang/pkg/usecase/banking"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	service := solid.NewSolidService(c.SolidEnv, c.SolidAPIKey)
	IBankingService := bankingService.NewBankingService(service)

	// HEALTH CHECK
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"msg": "Service running.",
		})
	})

	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		c.Set("Version", "v1")
		return c.Next()
	})

	person.RegisterRoutes(v1, service)
	bankingController.RegisterRoutes(v1, IBankingService)

	err = app.Listen(c.Port)
	if err != nil {
		log.Fatalln("Failed at Listen", err)
	}
}
