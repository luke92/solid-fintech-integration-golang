package contact

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
)

func (h handlerContact) GetContacts(c *fiber.Ctx) error {
	ctx := "get-contacts"
	personID := c.Params("personid")
	accountID := c.Query("accountid")

	response, err := h.service.GetContacts(personID, accountID)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
