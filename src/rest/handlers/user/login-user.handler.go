package user

import (
	"github.com/gofiber/fiber/v2"
	"goboilerplate.com/src/usecases"
	"goboilerplate.com/src/usecases/user"
)

type LoginUserHandler struct {
	loginUserUseCase user.ILoginUserUseCase
}

func NewLoginUserHandler(loginUserUseCase user.ILoginUserUseCase) *LoginUserHandler {
	return &LoginUserHandler{
		loginUserUseCase: loginUserUseCase,
	}
}

func (h *LoginUserHandler) LoginUser(c *fiber.Ctx) error {
	var req user.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	resp, err := h.loginUserUseCase.Apply(req)
	if err != nil {
		switch err {
		case usecases.ErrorUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
			})
		case usecases.ErrorInvalidCredentials:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid credentials",
			})
		case usecases.ErrorUserNotAbleToLogin:
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"message": "User is not able to login",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Internal server error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}