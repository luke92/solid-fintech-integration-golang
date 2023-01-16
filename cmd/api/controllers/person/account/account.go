package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
)

type handlerAccount struct {
	service solid.Service
}

func RegisterRoutes(app fiber.Router, service solid.Service) {
	h := &handlerAccount{
		service: service,
	}

	routes := app.Group("/account")

	routes.Post("/", h.AddAccount)
	routes.Get("/", h.GetAccounts)
}
