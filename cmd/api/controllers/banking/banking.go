package banking

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
)

type handlerBanking struct {
	service solid.Service
}

func RegisterRoutes(app fiber.Router, service solid.Service) {
	h := &handlerBanking{
		service: service,
	}

	routes := app.Group("/banking")

	routes.Post("/", h.AddFamily)
}
