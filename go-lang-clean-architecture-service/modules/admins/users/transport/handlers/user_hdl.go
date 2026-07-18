package handlers

import (
	"app/modules/admins/users/entity"
	"app/pkg/golangviet/templates"
	"app/views/admins/users"
	"app/views/pages"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type userUsc interface {
	FindOneUserUsc(id int) (*entity.User, error)
}

type userHdl struct {
	userUsc userUsc
}

func NewUserHdl(userUsc userUsc) *userHdl {
	return &userHdl{userUsc: userUsc}
}

func (u *userHdl) FindUserHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return templates.Render(c, users.Index())
	}
}

func (u *userHdl) UpdateUserViewHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")

		id, err := strconv.Atoi(idStr)

		if err != nil || id <= 0 {
			return templates.Render(c, pages.NotFound404("/admins/users/list"))
		}

		detailUser, err := u.userUsc.FindOneUserUsc(id)

		if err != nil || detailUser == nil {
			return templates.Render(c, pages.ErrorUserPage("Người dùng bạn tìm kiếm không tồn tại hoặc đã bị xóa!", "/admins/users/list"))
		}

		return templates.Render(c, users.UpdateUser(detailUser))
	}
}
