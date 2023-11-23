package usecases

import (
	"golang-project-template/internal/models"
)

type userUsecase struct {
	userRepository models.UserRepository
}

func NewUserUsecase(userRepository models.UserRepository) models.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) CreateUser(user *models.User) (int, error) {

	id, err := u.userRepository.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userUsecase) GetUserByID(userID int) (*models.User, error) {
	user, err := u.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetAllUsers() ([]*models.User, error) {

	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) UpdateUser(user *models.User) (*models.User, error) {
	user, err := u.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) DeleteUser(id int) error {
	err := u.userRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
