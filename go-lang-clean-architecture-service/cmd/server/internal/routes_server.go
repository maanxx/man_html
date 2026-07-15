package internal

import (
	"app/conf"
	"app/modules/admins/users"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

func RoutesServer(app *fiber.App, serviceCtx gosctx.ServiceContext) {
	users.SetupRoutesUser(app, serviceCtx)
	app.Static("/static", fmt.Sprintf("./%s", conf.UploadPathPublic))
}
