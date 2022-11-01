package usecase

import (
	"user-management-2/entity"
	"user-management-2/entity/response"
	"user-management-2/repository"

	"github.com/jinzhu/copier"
)

type IUserUsecase interface {
	CreateUser(user response.CreateUserRequest) error
	GetListUser() ([]response.GetUserResponse, error)
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (u UserUsecase) CreateUser(req response.CreateUserRequest) error {
	user := entity.User{}
	copier.Copy(&user, &req)

	err := u.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) GetList() ([]response.GetUserResponse, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, nil
	}
	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users)
	return userResponse, nil
}
