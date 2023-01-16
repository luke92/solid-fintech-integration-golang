package account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
)

func (h handlerAccount) GetAccounts(c *fiber.Ctx) error {
	personID := c.Params("personid")

	response, err := h.service.GetAccounts(personID)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   "add-account-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
