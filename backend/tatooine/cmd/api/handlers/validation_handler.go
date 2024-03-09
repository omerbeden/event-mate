package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/presenter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/validation/app/entrypoints"
)

func ValidateIdentity(service *entrypoints.ValidationService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody presenter.MernisRequest
		err := c.BodyParser(&requestBody)
		if err != nil {
			service.Logger.Infof("failed to parse body %s", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.BODY_PARSER_ERR,
			})
		}

		result, err := service.ValidateMernis(requestBody.NationalId, requestBody.Name, requestBody.LastName, requestBody.BirthYear)
		if err != nil {
			service.Logger.Info(err)
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.UNKNOW_ERR,
			})
		}

		return c.Status(fiber.StatusCreated).JSON(presenter.BaseResponse{
			APIVersion: presenter.APIVersion,
			Data:       result,
			Error:      "",
		})

	}
}
