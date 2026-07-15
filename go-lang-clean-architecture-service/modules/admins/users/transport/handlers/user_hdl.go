package handlers

import (
	"app/modules/admins/users/entity"
	"app/pkg/golangviet/templates"
	"app/views/admins/users"

	"github.com/gofiber/fiber/v2"
)

type UserUsc interface {
	InsertUserUsc(data *entity.User) (*int64, error)
	ListUserUsc() (*[]entity.User, error)
	FindUserByIDUsc(id int) (*entity.User, error)
	UpdateUserUsc(id int, data *entity.User) error
	DeleteUserUsc(id int) error
}

type userHdl struct {
	userUsc UserUsc
}

func NewUserHdl(userUsc UserUsc) *userHdl {
	return &userHdl{userUsc: userUsc}
}

func (h *userHdl) ListUserHdl() fiber.Handler {
	return func(c *fiber.Ctx) error{
		return templates.Render(c, users.Index())
	}
}
func (h *userHdl) ListUserAPIHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		draw := c.QueryInt("draw", 1)

		users, err := h.userUsc.ListUserUsc()

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		totalRecords := 0

		if users != nil {
			totalRecords = len(*users)
		}

		// return c.SendString("Hello Handler")
		// return templates.Render(c, users.Index())
		return c.JSON(fiber.Map{
			"draw":            draw,
			"recordsTotal":    totalRecords,
			"recordsFiltered": totalRecords,
			"data":            users,
		})
	}
}
