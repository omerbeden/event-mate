package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/presenter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/entrypoints"
)

func CreateActivity(service entrypoints.ActivityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody model.Activity
		err := c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.BODY_PARSER_ERR,
			})
		}

		res, err := service.CreateActivity(c.Context(), requestBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.UNKNOW_ERR,
			})
		}

		return c.Status(fiber.StatusCreated).JSON(presenter.BaseResponse{
			APIVersion: presenter.APIVersion,
			Data:       res,
			Error:      "",
		})
	}
}

func AddParticipant(service entrypoints.ActivityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		activityId, err := c.ParamsInt("activityId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.PARAM_PARSER_ERR,
			})
		}

		var requestBody model.User
		err = c.BodyParser(&requestBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.BODY_PARSER_ERR,
			})
		}

		if err := service.AddParticipant(requestBody, int64(activityId)); err != nil { // unnecessary int64 id , can be use int instead
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.UNKNOW_ERR,
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func GetParticipants(service entrypoints.ActivityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		activityId, err := c.ParamsInt("activityId")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.PARAM_PARSER_ERR,
			})
		}

		res, err := service.GetParticipants(int64(activityId))
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

func GetActivitiesByLocation(service entrypoints.ActivityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		city := c.Params("city")
		if city == "" {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.PARAM_PARSER_ERR,
			})
		}

		loc := model.Location{
			City: city,
		}

		res, err := service.GetActivitiesByLocation(c.Context(), loc)

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
