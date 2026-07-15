package repository

import (
	"app/modules/admins/users/entity"

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
func (s *userRepo) ListUserRepo() (*[]entity.User, error) {
	var users []entity.User

	result := s.db.Table(entity.User{}.TableName()).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

// read by id
func (s *userRepo) FindUserByID(id int) (*entity.User, error) {
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

	return result.Error
}

// delete
func (s *userRepo) DeleteUserRepo(id int) error {
	result := s.db.Table(entity.User{}.TableName()).Where("id=?", id).Delete(&entity.User{})

	return result.Error
}
