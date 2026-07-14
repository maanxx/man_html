package handlers

import (
	"app/modules/admins/users/entity"
	views_users "app/views/admins/users"

	"github.com/gofiber/fiber/v2"
)

type UserUsc interface {
	InsertUserUsc(data *entity.User) (*int64, error)
	ListUserUsc()([]entity.User, error)
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

func (h *userHdl) ListUserHdl() fiber.Handler{
	return func(c *fiber.Ctx) error {
		usersData, err := h.userUsc.ListUserUsc()

		if err != nil {
			return c.Status(500).SendString("Lỗi lấy dữ liệu từ DB")
		}

		c.Set("Content-Type", "text/html; charset=utf-8")

		component := views_users.UserListTemplate(usersData)

		return component.Render(c.Context(), c.Response().BodyWriter())
	}
}

func (h *userHdl) ShowCreateFormHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html; charset=utf-8")

		component := views_users.UserFormTemplate()

		return component.Render(c.Context(), c.Response().BodyWriter())
	}
}

func (h *userHdl) SubmitCreateFormHdl() fiber.Handler {
	return func(c *fiber.Ctx) error{
		var user entity.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).SendString("Dữ liệu không hợp lệ")
		}

		_, err := h.userUsc.InsertUserUsc(&user)

		if err != nil {
			return c.Status(500).SendString("Lỗi không lưu được vào DB")
		}

		return c.Redirect("/admins/users/list")
	}
}