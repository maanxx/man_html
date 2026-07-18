package repository

import (
	"app/modules/admins/users/entity"

	"github.com/teoit/gosctx/component/errs"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

// create
func (s *userRepo) InsertUserRepo(data *entity.User) (*int64, error) {
	result := s.db.Table(entity.User{}.TableName()).Create(data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data.ID, nil
}

// read
func (s *userRepo) FindUserRepo() (*[]entity.User, error) {
	var users []entity.User

	result := s.db.Table(entity.User{}.TableName()).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errs.ErrDataNotFound
	}

	return &users, nil
}

// read by id
func (s *userRepo) FindOneUserRepo(id int) (*entity.User, error) {
	var user entity.User

	result := s.db.Table(entity.User{}.TableName()).Where("id=?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// update
func (s *userRepo) UpdateUserRepo(id int, data *entity.User) error {
	result := s.db.Table(entity.User{}.TableName()).Where("id=?", id).Updates(data)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// delete
func (s *userRepo) DeleteUserRepo(ids []int) error {
	result := s.db.Table(entity.User{}.TableName()).Where("id IN ?", ids).Delete(&entity.User{})

	return result.Error
}
