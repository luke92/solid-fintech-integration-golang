package card

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
)

func (h handlerCard) CreatePinTokenCard(c *fiber.Ctx) error {
	ctx := "create-pin-token-card"
	personID := c.Params("personid")
	cardID := c.Params("cardid")

	response, err := h.service.CreatePinTokenCard(personID, cardID)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
