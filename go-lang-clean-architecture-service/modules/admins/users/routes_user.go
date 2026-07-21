package users

import (
	"app/modules/admins/users/composers"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

// views -> root (server) -> routes_server -> routes_user ->  handler(admins/users/list) || api(api/admins/users/create) -> usecase -> repo -> entity

//ORM: ánh xạ

func SetupRoutesUser(app *fiber.App, serviceCtx gosctx.ServiceContext) {

	group := app.Group("/admins/users")
	{
		fmt.Println(group)
		comp := composers.ComposerUserService(serviceCtx)
		group.Get("/list", comp.FindUserHdl()).Name("ecommerce.users.list")
		group.Get("/edit/:id", comp.UpdateUserViewHdl()).Name("ecommerce.users.edit")
	}

	groupApi := app.Group("/api/admins/users")
	{
		comp := composers.ComposerUserApiService(serviceCtx)
		groupApi.Post("/datatable", comp.FindUserApi()).Name("ecommerce.users.api.list")
		groupApi.Post("/create", comp.InsertUserApi())
		groupApi.Post("/edit", comp.UpdateUserApi())
		groupApi.Delete("/delete", comp.DeleteUserApi())
		groupApi.Patch("/update-status", comp.ChangeStatusApi())
	}
}
