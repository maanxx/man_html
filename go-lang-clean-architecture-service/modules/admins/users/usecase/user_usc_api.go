package usecase

import (
	"app/modules/admins/users/entity"
	"app/modules/admins/users/mapping"
	"app/modules/admins/users/transport/requests"
	"app/modules/admins/users/transport/respones"
	"app/pkg/golangviet/utils"
	"context"
	"fmt"
	"strings"

	"app/modules/admins/common/errs"
)

// repo -> usc
type userApi interface {
	FindUserDatatableRepo(ctx context.Context, filter *utils.Filters) ([]*entity.User, error)
	FindUserRepo() (*[]entity.User, error)
	FindOneUserRepo(id int) (*entity.User, error)
	InsertUserRepo(data *entity.User) (*int64, error)
	UpdateUserRepo(id int, data *entity.User) error
	UpdateStatusRepo(id int, status int) error
	DeleteUserRepo(ids []int) error
}

type userApiUsc struct {
	userApi userApi
}

func NewUserApiUsc(userApi userApi) *userApiUsc {
	return &userApiUsc{
		userApi: userApi,
	}
}

// create
func (u *userApiUsc) InsertUserApiUsc(req *requests.UserCreation) (*int64, error) {
	userEntity := &entity.User{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		Email:     req.Email,
		Status:    req.Status,
	}

	newUser, err := u.userApi.InsertUserRepo(userEntity)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, errs.ErrCreateUserFailed
		}
		return nil, fmt.Errorf("InsertUserApiUsc: %w", err)
	}

	return newUser, nil
}

// read
func (u *userApiUsc) FindUserApiUsc() (*[]respones.DataTableUserResp, error) {
	_, err := u.userApi.FindUserRepo()

	if err != nil {
		if err == errs.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return nil, nil
}

// read
func (u *userApiUsc) FindListDatatableUserApiUsc(ctx context.Context, req *requests.ListDatatableUserReq) (*respones.ListDatatableUserResp, error) {
	limit := req.Length

	conds := mapping.MapperCondListDatatableUser(req)

	filter := utils.Filters{
		Conds:    &conds,
		Offset:   req.Start,
		PageSize: limit,
	}

	data, err := u.userApi.FindUserDatatableRepo(ctx, &filter)

	if err != nil {
		if err == errs.ErrInternalServer {
			return nil, nil
		}
		return nil, err
	}

	result := mapping.MapperDatatable(data)

	userDatatable := &respones.ListDatatableUserResp{
		Draw:            req.Draw,
		RecordsTotal:    20,
		RecordsFiltered: 10,
		Data:            result,
	}

	return userDatatable, nil
}

// read by id
func (u *userApiUsc) FindOneUserApiUsc(id int) (*entity.User, error) {
	data, err := u.userApi.FindOneUserRepo(id)

	if err != nil {
		if err == errs.ErrInternalServer {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

// update
func (u *userApiUsc) UpdateUserApiUsc(id int, req *requests.PayloadUserCreation) error {

	user := entity.User{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		Email:     req.Email,
		Status:    req.Status,
	}

	err := u.userApi.UpdateUserRepo(id, &user)

	if err != nil {
		return errs.ErrDataNotFound
	}

	return nil
}

// udpate status
func (u *userApiUsc) ChangeStatusApiUsc(id int, status int) error {
	err := u.userApi.UpdateStatusRepo(id, status)

	if err != nil {
		return errs.ErrDataNotFound
	}

	return nil
}

// delete
func (u *userApiUsc) DeleteUserApiUsc(ids []int) error {
	err := u.userApi.DeleteUserRepo(ids)

	if err != nil {
		return errs.ErrDataNotFound
	}

	return nil
}
