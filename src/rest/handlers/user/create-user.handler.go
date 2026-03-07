package user

import (
	"github.com/gofiber/fiber/v2"
	"goboilerplate.com/src/usecases"
	"goboilerplate.com/src/usecases/user"
)

type CreateUserHandler struct {
	createUserUseCase user.ICreateUserUseCase
}

func NewCreateUserHandler(createUserUseCase user.ICreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (h *CreateUserHandler) CreateUser(c *fiber.Ctx) error {
	var req user.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	resp, err := h.createUserUseCase.Apply(req)
	if err != nil {
		switch err {
		case usecases.ErrUserAlreadyExists:
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  "error",
				"message": "User already exists",
			})
		case usecases.ErrCannotCreateUser:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Cannot create user",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Internal server error",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}