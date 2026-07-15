package usecase

import "app/modules/admins/usergroups/entity"

type UsergroupRepo interface {
	ListUserGroupRepo() (*[]entity.UserGroup, error)
}

type usergroupUseCase struct {
	usergroupRepo UsergroupRepo
}

func NewUserGroupUsc(usergroupRepo UsergroupRepo) *usergroupUseCase {
	return &usergroupUseCase{usergroupRepo: usergroupRepo}
}

func (u *usergroupUseCase)ListUserGroupUsc() (*[]entity.UserGroup, error) {
	return u.usergroupRepo.ListUserGroupRepo()
}