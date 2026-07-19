package composers

import (
	"app/modules/admins/users/entity"
	"app/modules/admins/users/repository"
	"app/modules/admins/users/transport/api"
	"app/modules/admins/users/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/configs"
)

type composerUserApi interface {
	FindUserApi() fiber.Handler
	InsertUserApi() fiber.Handler
	UpdateUserApi() fiber.Handler
	DeleteUserApi() fiber.Handler
}

func ComposerUserApiService(serviceCtx gosctx.ServiceContext) composerUserApi {
	db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	db.AutoMigrate(&entity.User{})

	repo := repository.NewUserRepo(db)

	uscApi := usecase.NewUserApiUsc(repo)

	api := api.NewUserApi(uscApi)

	return api
}
