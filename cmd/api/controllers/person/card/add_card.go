package card

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (h handlerCard) AddCard(c *fiber.Ctx) error {
	ctx := "add-card"
	personID := c.Params("personid")
	body := models.NewCard{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-parse-request",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	response, err := h.service.AddCard(personID, body)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
