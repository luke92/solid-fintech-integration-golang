package contact

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
)

type handlerContact struct {
	service solid.Service
}

func RegisterRoutes(app fiber.Router, service solid.Service) {
	h := &handlerContact{
		service: service,
	}

	routes := app.Group("/contact")

	routes.Post("/", h.AddContact)
	routes.Get("/", h.GetContacts)
}
