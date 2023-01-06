package family

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
)

type handlerFamily struct {
	service solid.Service
}

func RegisterRoutes(app fiber.Router, service solid.Service) {
	h := &handlerFamily{
		service: service,
	}

	routes := app.Group("/family")

	routes.Post("/", h.AddFamily)
}
