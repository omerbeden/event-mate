package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/omerbeden/event-mate/backend/tatooine/cmd/api/presenter"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/domain/model"
	"github.com/omerbeden/event-mate/backend/tatooine/modules/profile/app/entrypoints"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg"
	customerrors "github.com/omerbeden/event-mate/backend/tatooine/pkg/customErrors"
	"go.uber.org/zap"
)

func CreateUserProfile(service entrypoints.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody model.UserProfile
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

		err = service.CreateUser(ctx, &requestBody)
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
			Data:       "ok",
			Error:      "",
		})
	}
}

func GetCurrentUserProfile(service entrypoints.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		externalId := c.Params("externalId")

		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()

		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}

		newLogger := service.Logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		res, err := service.GetCurrentUserProfile(ctx, externalId)

		if err != nil {
			service.Logger.Error(err)
			if err == customerrors.ERR_NOT_FOUND {
				return c.Status(fiber.StatusNotFound).JSON(presenter.BaseResponse{
					APIVersion: presenter.APIVersion,
					Data:       nil,
					Error:      customerrors.ERR_NOT_FOUND.Error(),
				})
			}

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

func UpdateProfileImageUrl(service entrypoints.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(presenter.ProfileImageUpdateRequest)

		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()
		externalId := c.Params("externalId")

		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}

		newLogger := service.Logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		err = service.UpdateProfileImage(ctx, externalId, request.ProfileImageUrl)
		if err != nil {
			service.Logger.Error(err)
			if err == customerrors.ERR_NOT_FOUND {
				return c.Status(fiber.StatusNotFound).JSON(presenter.BaseResponse{
					APIVersion: presenter.APIVersion,
					Data:       nil,
					Error:      customerrors.ERR_NOT_FOUND.Error(),
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.UNKNOW_ERR,
			})
		}

		return c.Status(fiber.StatusOK).JSON(presenter.BaseResponse{
			APIVersion: presenter.APIVersion,
			Data:       "OK",
			Error:      "",
		})
	}
}
