package card

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
)

type handlerCard struct {
	service solid.Service
}

func RegisterRoutes(app fiber.Router, service solid.Service) {
	h := &handlerCard{
		service: service,
	}

	routes := app.Group("/card")

	routes.Post("/", h.AddCard)
	routes.Patch("/:cardid/activate", h.ActivateCard)
	routes.Post("/:cardid/pintoken", h.CreatePinTokenCard)
	routes.Post("/:cardid/showtoken", h.CreateShowTokenCard)
}
