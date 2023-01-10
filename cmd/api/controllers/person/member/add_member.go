package member

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
	"github.com/luke92/solid-fintech-integration-golang/pkg/common/models"
)

func (h handlerMember) AddMember(c *fiber.Ctx) error {
	personID := c.Params("personid")
	body := models.NewMemberData{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   "add-member-parse-request",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	response, err := h.service.AddMember(personID, body)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   "add-member-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
