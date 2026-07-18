package usecase

import (
	"app/modules/admins/users/entity"
)

// repo -> usc
type UserRepo interface {
	// InsertUserRepo(data *entity.User) (*int64, error)
	FindUserRepo() (*[]entity.User, error)
	FindOneUserRepo(id int) (*entity.User, error)
	// UpdateUserRepo(id int, data *entity.User) error
	// DeleteUserRepo(id int) error
}

type userUseCase struct {
	userRepo UserRepo
}

func NewUserUsc(userRepo UserRepo) *userUseCase {
	return &userUseCase{userRepo: userRepo}
}

// create
// func (u *userUseCase) InsertUserUsc(data *entity.User) (*int64, error) {
// 	return u.userRepo.InsertUserRepo(data)
// }

// read
func (u *userUseCase) FindUserUsc() (*[]entity.User, error) {
	data, err := u.userRepo.FindUserRepo()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// read by id
func (u *userUseCase) FindOneUserUsc(id int) (*entity.User, error) {
	return u.userRepo.FindOneUserRepo(id)
}

// // update
// func (u *userUseCase) UpdateUserUsc(id int, data *entity.User) error {
// 	return u.userRepo.UpdateUserRepo(id, data)
// }

// // delete
// func (u *userUseCase) DeleteUserUsc(id int) error {
// 	return u.userRepo.DeleteUserRepo(id)
// }
