package person

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
)

func (h handler) SubmitKYC(c *fiber.Ctx) error {
	ctx := "submit-kyc-handler"
	personID := c.Params("personid")

	response, err := h.service.SubmitKYC(personID)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
