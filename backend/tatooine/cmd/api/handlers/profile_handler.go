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
		logger := pkg.Logger()

		var requestBody model.CreateUserProfileRequest
		err := c.BodyParser(&requestBody)
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

		err = service.CreateUser(ctx, &requestBody)
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
			Data:       true,
			Error:      "",
		})
	}
}

func GetCurrentUserProfile(service entrypoints.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := pkg.Logger()
		externalId := c.Params("externalId")

		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()

		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}

		newLogger := logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		profile, err := service.GetCurrentUserProfile(ctx, externalId)

		if err != nil {
			logger.Error(err)
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

		result := presenter.ProfileToGetUserResponse(*profile)

		return c.Status(fiber.StatusOK).JSON(presenter.BaseResponse{
			APIVersion: presenter.APIVersion,
			Data:       result,
			Error:      "",
		})
	}
}

func UpdateProfileImageUrl(service entrypoints.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := pkg.Logger()
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

		newLogger := logger.With(zap.String("requestid", requestid))
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		err = service.UpdateProfileImage(ctx, externalId, request.ProfileImageUrl)
		if err != nil {
			logger.Error(err)
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

func UpdateProfileVerification(service entrypoints.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := pkg.Logger()
		request := new(presenter.ProfileVerificationUpdateRequest)

		requestid, err := getRequestId(c)
		if err != nil {
			return err
		}
		newLogger := logger.With(zap.String("requestid", requestid))
		ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
		defer cancel()
		ctx = context.WithValue(ctx, pkg.LoggerKey, newLogger)

		if err := c.BodyParser(request); err != nil {
			logger.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       nil,
				Error:      presenter.BODY_PARSER_ERR,
			})
		}

		externalId := c.Params("externalId")

		err = service.UpdateVerification(ctx, request.IsVerified, externalId)
		if err != nil {
			logger.Error(err)
			if err == customerrors.ERR_NOT_FOUND {
				return c.Status(fiber.StatusNotFound).JSON(presenter.BaseResponse{
					APIVersion: presenter.APIVersion,
					Data:       false,
					Error:      customerrors.ERR_NOT_FOUND.Error(),
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.BaseResponse{
				APIVersion: presenter.APIVersion,
				Data:       false,
				Error:      presenter.UNKNOW_ERR,
			})
		}

		return c.Status(fiber.StatusOK).JSON(presenter.BaseResponse{
			APIVersion: presenter.APIVersion,
			Data:       true,
			Error:      "",
		})
	}
}
