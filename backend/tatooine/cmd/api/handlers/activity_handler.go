package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/presenter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/activity/app/entrypoints"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	customerrors "github.com/omerbeden/event-mate/backend/tatooine/pkg/customErrors"
	"go.uber.org/zap"
)

func getRequestId(c *fiber.Ctx) (string, error) {
	requestid, ok := c.Locals("requestid").(string)

	if !ok {
		return "", c.Status(fiber.StatusInternalServerError).JSON(presenter.BaseResponse{
			APIVersion: presenter.APIVersion,
			Data:       "reqeust id ",
			Error:      presenter.UNKNOW_ERR,
		})
	}
	return requestid, nil
}

func CreateActivity(service entrypoints.ActivityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody model.Activity
		err := c.BodyParser(&requestBody)
		if err != nil {
			service.Logger.Error(presenter.BODY_PARSER_ERR)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.BODY_PARSER_ERR,
			})
		}

		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()

		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}

		newLogger := service.Logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		res, err := service.CreateActivity(ctx, requestBody)
		if err != nil {
			service.Logger.Error(err)

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
			service.Logger.Error(presenter.PARAM_PARSER_ERR)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.PARAM_PARSER_ERR,
			})
		}

		var requestBody model.User
		err = c.BodyParser(&requestBody)
		if err != nil {
			service.Logger.Error(presenter.BODY_PARSER_ERR)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.BODY_PARSER_ERR,
			})
		}

		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()

		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}

		newLogger := service.Logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		if err := service.AddParticipant(ctx, requestBody, int64(activityId)); err != nil { // unnecessary int64 id , can be use int instead
			service.Logger.Error(err)
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
			service.Logger.Error(presenter.PARAM_PARSER_ERR)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.PARAM_PARSER_ERR,
			})
		}

		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()

		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}

		newLogger := service.Logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)
		res, err := service.GetParticipants(ctx, int64(activityId))
		if err != nil {
			service.Logger.Error(err)
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
		city := c.Query("city")
		if city == "" {
			service.Logger.Error(presenter.PARAM_PARSER_ERR)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.PARAM_PARSER_ERR,
			})
		}

		loc := model.Location{
			City: city,
		}

		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()

		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}

		newLogger := service.Logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		res, err := service.GetActivitiesByLocation(ctx, loc)

		if err != nil && err != customerrors.ErrActivityDoesNotHaveParticipants {
			service.Logger.Error(err)
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

func GetActivityById(service entrypoints.ActivityService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		activityId := c.Params("activityId")

		aI, err := strconv.Atoi(activityId)

		if err != nil {
			service.Logger.Error(presenter.PARAM_PARSER_ERR)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.PARAM_PARSER_ERR,
			})
		}

		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()

		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}

		newLogger := service.Logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		res, err := service.GetActivityById(ctx, int64(aI))

		if err != nil {
			service.Logger.Error(err)
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
