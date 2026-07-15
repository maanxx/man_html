package composers

import (
	"app/modules/admins/usergroups/handler"
	"app/modules/admins/usergroups/repository"
	"app/modules/admins/usergroups/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/configs"
)

type composerUserGroupHdl interface {
	ListUserGroupHdl() fiber.Handler
}

func ComposerUserGroupService(serviceCtx gosctx.ServiceContext) composerUserGroupHdl {
	db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	repo := repository.NewUserGroupRepo(db)

	usc := usecase.NewUserGroupUsc(repo)

	hdl := handler.NewUserGroupHdl(usc)

	return hdl

}