package banking

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (h handlerBanking) AddFamily(c *fiber.Ctx) error {
	body := models.PersonDataPartial{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   "add-person-parse-request",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	response, err := h.service.AddPerson(body)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   "add-person-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
