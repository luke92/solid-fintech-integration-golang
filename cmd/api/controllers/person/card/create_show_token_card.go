package card

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
)

func (h handlerCard) CreateShowTokenCard(c *fiber.Ctx) error {
	ctx := "create-show-token-card"
	personID := c.Params("personid")
	cardID := c.Params("cardid")

	response, err := h.service.CreateShowTokenCard(personID, cardID)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
