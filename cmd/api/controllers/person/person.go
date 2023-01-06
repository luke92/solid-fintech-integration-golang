package person

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
)

type handler struct {
	service solid.Service
}

func RegisterRoutes(app fiber.Router, service solid.Service) {
	h := &handler{
		service: service,
	}

	routes := app.Group("/person")

	routes.Post("/", h.AddPerson)
	routes.Post("/:personid/kyc", h.SubmitKYC)
	routes.Post("/:personid/idv", h.SubmitIDV)
	//routes.Get("/:id", h.GetBook)
	//routes.Put("/:id", h.UpdateBook)
}
