package handler

import (
	"app/modules/admins/usergroups/entity"

	"github.com/gofiber/fiber/v2"
)

type UsergroupUcs interface {
	ListUserGroupUsc() (*[]entity.UserGroup, error)
}

type usergroupHdl struct {
	usergroupUcs UsergroupUcs
}

func NewUserGroupHdl(usergroupUcs UsergroupUcs) *usergroupHdl {
	return &usergroupHdl{usergroupUcs: usergroupUcs}
}

func (u *usergroupHdl) ListUserGroupHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Run handler")
	}
}
