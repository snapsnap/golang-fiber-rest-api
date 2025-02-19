package api

import (
	"api-dev/domain"
	"api-dev/dto"
	"api-dev/internal/middleware"
	"api-dev/internal/utils"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type userApi struct {
	userService domain.UserService
}

func NewUser(app *fiber.App, userService domain.UserService) {
	ua := userApi{
		userService: userService,
	}

	app.Post("/register", ua.Register)

	userGroup := app.Group("/users")
	userGroup.Use(middleware.JWTMiddleware())
	userGroup.Post("/", ua.Index)
}

func (ua userApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	req := dto.UserRequest{
		Limit: 10,
		Page:  1,
	}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(dto.CreateResponseError("Invalid request body", ""))
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}

	res, err := ua.userService.Index(c, req.Limit, req.Page)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error(), ""))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ua userApi) Register(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.RegisterUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := utils.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("validation failed", fails))
	}
	// Hash password sebelum menyimpan ke database
	hashedPassword, errHash := utils.HashPassword(req.Password)
	if errHash != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError("failed to hash password", ""))
	}
	req.Password = hashedPassword

	err := ua.userService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error(), ""))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
