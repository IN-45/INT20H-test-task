package handler

import (
	user_service "github.com/IN-45/INT20H-test-task/modules/user/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	validator   *validator.Validate
	authService *user_service.AuthService
}

func NewAuthHandler(
	validator *validator.Validate,
	authService *user_service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		validator:   validator,
		authService: authService,
	}
}

func RegisterAuthHandler(app *fiber.App, h *AuthHandler) {
	app.Post("/sign-in", h.SignIn)
	app.Post("/sign-up", h.SignUp)
}

// SignIn
//
//	@Summary	Sign in to account
//	@Tags		Authentication
//	@Accept		json
//	@Param		input	body	dtoAuth	true	"User credentials"
//	@Success	200		"Saves token to cookie"
//	@Failure	400		{object}	fiber.Error
//	@Failure	500		{object}	fiber.Error
//	@Router		/sign-in [post]
func (h *AuthHandler) SignIn(ctx *fiber.Ctx) error {
	dto := new(dtoAuth)

	if err := ctx.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	jwtToken, err := h.authService.SignIn(ctx.Context(), dto.Email, dto.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	id, err := h.authService.GetUserIDByEmail(ctx.Context(), dto.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: jwtToken,
	})
	ctx.Cookie(&fiber.Cookie{
		Name:  "user_id",
		Value: id,
		// Secure:   true,
		// HTTPOnly: true,
	})

	return nil
}

// SignUp
//
//	@Summary	Sign up into account
//	@Tags		Authentication
//	@Accept		json
//	@Param		input	body	dtoAuth	true	"User credentials"
//	@Success	200
//	@Failure	400	{object}	fiber.Error
//	@Failure	500	{object}	fiber.Error
//	@Router		/sign-up [post]
func (h *AuthHandler) SignUp(ctx *fiber.Ctx) error {
	dto := new(dtoAuth)

	if err := ctx.BodyParser(dto); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.validator.Struct(dto); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return h.authService.SignUp(ctx.Context(), dto.Email, dto.Password)
}

type dtoAuth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}
