package usergroups

import (
	"app/modules/admins/usergroups/composers"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

func SetupRoutesUserGroup(app *fiber.App, serviceCtx gosctx.ServiceContext) {
	group := app.Group("/admins/usergroups")
	{
		comp := composers.ComposerUserGroupService(serviceCtx)
		group.Get("/list", comp.ListUserGroupHdl()).Name("ecommerce.users.list")
	}
}