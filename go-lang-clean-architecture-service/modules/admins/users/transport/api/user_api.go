package api

import (
	"app/modules/admins/users/transport/requests"
	"app/modules/admins/users/transport/respones"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx/core"
)

type userApiUsc interface {
	FindUserApiUsc() (*[]respones.DataTableUserResp, error)
	FindListDatatableUserApiUsc(ctx context.Context, req *requests.ListDatatableUserReq) (*respones.ListDatatableUserResp, error)
	InsertUserApiUsc(req *requests.UserCreation) (*int64, error)
	UpdateUserApiUsc(id int, req *requests.PayloadUserCreation) error
	ChangeStatusApiUsc(id int, status int) error
	DeleteUserApiUsc(ids []int) error
}

type userApi struct {
	userApiUsc userApiUsc
}

func NewUserApi(userApiUsc userApiUsc) *userApi {
	return &userApi{userApiUsc: userApiUsc}
}

// create
func (u *userApi) InsertUserApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UserCreation

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Data create user not found",
			})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		_, err := u.userApiUsc.InsertUserApiUsc(&req)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Insert Successfully",
		})
	}
}

// read
// func (u *userApi) FindUserApi() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		users, err := u.userApiUsc.FindUserApiUsc()
// 		if err != nil {
// 			return c.JSON(fiber.Map{
// 				"data": "",
// 			})
// 		}

// 		return c.JSON(fiber.Map{
// 			"data": users,
// 		})
// 	}
// }

func (u *userApi) FindUserApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.ListDatatableUserReq

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data!",
			})
		}

		resp, err := u.userApiUsc.FindListDatatableUserApiUsc(c.Context(), &req)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Server error",
			})
		}
		return c.JSON(resp)
	}
}

// update
func (u *userApi) UpdateUserApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.PayloadUserCreation

		if err := c.BodyParser(&req); err != nil {
			return core.ReturnErrsForApi(c, err)
		}

		if req.ID <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		errUsc := u.userApiUsc.UpdateUserApiUsc(req.ID, &req)

		if errUsc != nil {
			return core.ReturnErrsForApi(c, errUsc)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  1,
			"message": "Update Successfully",
		})
	}
}

// update status
func (u *userApi) ChangeStatusApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.PayloadChangeStatus

		if err := c.BodyParser(&req); err != nil {
			return core.ReturnErrsForApi(c, err)
		}

		if req.ID <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		errUsc := u.userApiUsc.ChangeStatusApiUsc(req.ID, req.Status)

		if errUsc != nil {
			return core.ReturnErrsForApi(c, errUsc)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  1,
			"message": "Update Successfully",
		})
	}
}

func (u *userApi) DeleteUserApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			IDs []int `json:"ids"`
		}

		if err := c.BodyParser(&req); err != nil {
			return core.ReturnErrsForApi(c, err)
		}

		if len(req.IDs) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"err": "Invalid user IDs",
			})
		}

		errUsc := u.userApiUsc.DeleteUserApiUsc(req.IDs)

		if errUsc != nil {
			return core.ReturnErrsForApi(c, errUsc)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  1,
			"message": "Delete successfully",
		})

	}
}
