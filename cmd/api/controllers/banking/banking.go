package banking

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/pkg/usecase"
)

type handlerBanking struct {
	bankingService usecase.IBankingService
}

func RegisterRoutes(app fiber.Router, service usecase.IBankingService) {
	h := &handlerBanking{
		bankingService: service,
	}

	routes := app.Group("/banking")

	routes.Post("/", h.AddFamily)
}
