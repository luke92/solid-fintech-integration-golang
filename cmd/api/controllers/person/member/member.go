package member

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
)

type handlerMember struct {
	service solid.Service
}

func RegisterRoutes(app fiber.Router, service solid.Service) {
	h := &handlerMember{
		service: service,
	}

	routes := app.Group("/member")

	routes.Post("/", h.AddMember)
}
