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
		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()

		logger := pkg.Logger()
		newLogger := logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		var requestBody presenter.CreateActivityRequest
		err = c.BodyParser(&requestBody)
		if err != nil {
			newLogger.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.BODY_PARSER_ERR,
			})
		}

		activity := toActivity(requestBody)
		res, err := service.CreateActivity(ctx, activity)
		if err != nil {
			logger.Error(err)

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
func toActivity(request presenter.CreateActivityRequest) model.Activity {
	return model.Activity{
		Title:    request.Title,
		Category: request.Category,
		CreatedBy: model.User{
			ID:         request.CreatedById,
			ExternalId: request.CreatedByExternalId,
		},
		Location: model.Location{
			City:        request.Location.Location.City,
			District:    request.Location.Location.District,
			Description: request.Location.Description,
			Latitude:    request.Location.Latitude,
			Longitude:   request.Location.Longitude,
		},
		StartAt:           time.Time{},
		EndAt:             time.Time{},
		Content:           request.Content,
		Rules:             request.Rules,
		Flow:              request.Flow,
		Quota:             request.Quota,
		GenderComposition: model.GenderComposition(request.GenderComposition),
	}
}

func AddParticipant(service entrypoints.ActivityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := pkg.Logger()
		activityId, err := c.ParamsInt("activityId")
		if err != nil {
			logger.Error(presenter.PARAM_PARSER_ERR)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.PARAM_PARSER_ERR,
			})
		}

		var requestBody model.User
		err = c.BodyParser(&requestBody)
		if err != nil {
			logger.Error(presenter.BODY_PARSER_ERR)
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

		newLogger := logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		if err := service.AddParticipant(ctx, requestBody, int64(activityId)); err != nil { // unnecessary int64 id , can be use int instead
			logger.Error(err)
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
		logger := pkg.Logger()
		activityId, err := c.ParamsInt("activityId")
		if err != nil {
			logger.Error(presenter.PARAM_PARSER_ERR)
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

		newLogger := logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)
		res, err := service.GetParticipants(ctx, int64(activityId))
		if err != nil {
			logger.Error(err)
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
		logger := pkg.Logger()
		city := c.Query("city")
		if city == "" {
			logger.Error(presenter.PARAM_PARSER_ERR)
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

		newLogger := logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		res, err := service.GetActivitiesByLocation(ctx, loc)

		if err != nil && err != customerrors.ErrActivityDoesNotHaveParticipants {
			logger.Error(err)
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
		logger := pkg.Logger()
		activityId := c.Params("activityId")

		aI, err := strconv.Atoi(activityId)

		if err != nil {
			logger.Error(presenter.PARAM_PARSER_ERR)
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

		newLogger := logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		res, err := service.GetActivityById(ctx, int64(aI))

		if err != nil {
			logger.Error(err)
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
