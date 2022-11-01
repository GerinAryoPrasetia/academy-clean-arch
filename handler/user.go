package handler

import (
	"user-management-2/entity/response"
	"user-management-2/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(usercase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: usercase}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := response.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.CreateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to create user",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "user created successfully",
		Data:    nil,
	})
}

func (h UserHandler) GetList(ctx *fiber.Ctx) error {
	users, err := h.userUsecase.GetList()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "get list users failed",
			Data:    nil,
		})
	}

	if len(users) < 0 {
		return ctx.Status(fiber.StatusNoContent).JSON(response.BaseResponse{
			Code:    fiber.StatusNoContent,
			Message: "no user found",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "successfully get all users",
		Data:    users,
	})
}
