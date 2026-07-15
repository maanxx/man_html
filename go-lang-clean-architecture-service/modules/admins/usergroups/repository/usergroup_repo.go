package repository

import (
	"app/modules/admins/usergroups/entity"

	"gorm.io/gorm"
)

type usergroupRepo struct {
	db *gorm.DB
}

func NewUserGroupRepo(db *gorm.DB) *usergroupRepo {
	return &usergroupRepo{db: db}
}

func (u *usergroupRepo)ListUserGroupRepo() (*[]entity.UserGroup, error) {
	var usergroups []entity.UserGroup

	var result = u.db.Table(entity.UserGroup{}.TableName()).Find(&usergroups)

	if result != nil {
		return nil, result.Error
	}

	return &usergroups, nil
}