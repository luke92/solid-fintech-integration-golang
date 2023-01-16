package banking

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luke92/solid-fintech-integration-golang/cmd/api/controllers"
	"github.com/luke92/solid-fintech-integration-golang/pkg/types"
)

func (h handlerBanking) AddFamily(c *fiber.Ctx) error {
	ctx := "add-family"
	input := types.AddPersonAndFamilyInput{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&input); err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-parse-request",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(errResponse)
	}

	response, err := h.bankingService.AddPersonAndFamily(input)
	if err != nil {
		errResponse := controllers.ErrorResponse{
			Title:   ctx + "-service",
			Message: err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
