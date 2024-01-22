package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/presenter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/entrypoints"
)

func CreateUserProfile(service entrypoints.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody model.UserProfile
		err := c.BodyParser(&requestBody)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.BODY_PARSER_ERR,
			})
		}

		err = service.CreateUser(&requestBody)
		if err != nil {
			fmt.Printf("err: %v\n", err)

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.UNKNOW_ERR,
			})
		}

		return c.Status(fiber.StatusCreated).JSON(presenter.BaseResponse{
			APIVersion: presenter.APIVersion,
			Data:       "ok",
			Error:      "",
		})
	}
}

func GetCurrentUserProfile(service entrypoints.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		externalId := c.Params("externalId")

		res, err := service.GetCurrentUserProfile(externalId)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.UNKNOW_ERR,
			})
		}

		return c.Status(fiber.StatusOK).JSON(presenter.BaseResponse{
			APIVersion: presenter.APIVersion,
			Data:       res,
			Error:      "",
		})
	}
}
