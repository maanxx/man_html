package internal

import (
	"app/modules/admins/users"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

func RoutesServer(app *fiber.App, serviceCtx gosctx.ServiceContext) {
	users.SetupRoutesUser(app, serviceCtx)
}

