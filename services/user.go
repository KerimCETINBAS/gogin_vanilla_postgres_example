package services

import (
	"github.com/kerimcetinbas/goginpostgrestut/repositories"
	"github.com/kerimcetinbas/goginpostgrestut/types"
)

type userService struct {
	userRepository repositories.IUserRepository
}
type IUserService interface {
	CreateUser(data *types.UserCreateDto) error
	GetUsers() (*[]types.User, error)
}

func UserService() IUserService {
	return &userService{
		userRepository: repositories.UserRepository(),
	}
}

func (s *userService) CreateUser(data *types.UserCreateDto) error {
	return s.userRepository.CreateUser(data)
}

func (s *userService) GetUsers() (*[]types.User, error) {
	return s.userRepository.GetUsers()
}
