package users

import (
	"app/modules/admins/users/composers"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

// views -> root (server) -> routes_server -> routes_user ->  handler(admins/users/list) || api(api/admins/users/create) -> usecase -> repo -> entity

//ORM: ánh xạ

func SetupRoutesUser(app *fiber.App, serviceCtx gosctx.ServiceContext) {
	group := app.Group("/admins/users")
	{
		comp := composers.ComposerUserService(serviceCtx)
		group.Get("/list", comp.ListUserHdl()).Name("ecommerce.users.list")
		group.Get("/create", comp.ShowCreateFormHdl())
	}

	groupApi := app.Group("/admins/users")
	{
		comp := composers.ComposerUserService(serviceCtx)
		groupApi.Post("/create", comp.SubmitCreateFormHdl())
	}
	
}