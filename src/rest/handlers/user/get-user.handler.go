package user

import (
	"github.com/gofiber/fiber/v2"
	"goboilerplate.com/src/usecases"
	"goboilerplate.com/src/usecases/user"
)

type GetUserHandler struct {
	getUserUseCase user.IGetUserUseCase
}

func NewGetUserHandler(getUserUseCase user.IGetUserUseCase) *GetUserHandler {
	return &GetUserHandler{
		getUserUseCase: getUserUseCase,
	}
}

func (h *GetUserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	resp, err := h.getUserUseCase.Apply(userID)
	if err != nil {
		switch err {
		case usecases.ErrUserNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
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