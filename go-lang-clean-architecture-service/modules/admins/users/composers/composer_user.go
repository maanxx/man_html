package composers

import (
	"app/modules/admins/users/entity"
	"app/modules/admins/users/handlers"
	"app/modules/admins/users/repository"
	"app/modules/admins/users/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/configs"
)

type composerUserHdl interface {
	ListUserHdl() fiber.Handler
	ShowCreateFormHdl() fiber.Handler
	SubmitCreateFormHdl() fiber.Handler	
}

func ComposerUserService(serviceCtx gosctx.ServiceContext) composerUserHdl {
	db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	db.AutoMigrate(&entity.User{})

	repo := repository.NewUserRepo(db)

	usc := usecase.NewUserRepo(repo)

	hdl := handlers.NewUserHdl(usc)

	return hdl
}