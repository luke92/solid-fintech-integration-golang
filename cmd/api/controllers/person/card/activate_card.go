package card

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (h handlerCard) ActivateCard(c *fiber.Ctx) error {
	ctx := "activate-card"
	personID := c.Params("personid")
	cardID := c.Params("cardid")
	body := models.ExpireAndLast4CardData{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-parse-request",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	response, err := h.service.ActivateCard(personID, cardID, body)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
