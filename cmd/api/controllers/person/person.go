package person

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers/person/account"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers/person/family"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers/person/member"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/services/solid"
)

type handlerPerson struct {
	service solid.Service
}

func RegisterRoutes(app fiber.Router, service solid.Service) {
	h := &handlerPerson{
		service: service,
	}

	routes := app.Group("/person")

	routes.Post("/", h.AddPerson)

	subRoutePersonID := routes.Group("/:personid")
	subRoutePersonID.Post("/kyc", h.SubmitKYC)
	subRoutePersonID.Post("/idv", h.SubmitIDV)

	family.RegisterRoutes(subRoutePersonID, service)
	account.RegisterRoutes(subRoutePersonID, service)
	member.RegisterRoutes(subRoutePersonID, service)
}
