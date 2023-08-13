package services

import (
	"errors"

	"github.com/kerimcetinbas/goginpostgrestut/repositories"
	"github.com/kerimcetinbas/goginpostgrestut/types"
)

type authService struct {
	userRepository repositories.IUserRepository
}

type IAuthService interface {
	Login(data *types.UserLoginDto) (types.User, error)
}

func AuthService() IAuthService {
	return &authService{
		userRepository: repositories.UserRepository(),
	}
}

func (s *authService) Login(data *types.UserLoginDto) (types.User, error) {

	var (
		user types.User
		err  error
	)
	user, err = s.userRepository.FindUserByName(data.Name)

	if user.Password != data.Password {
		return user, errors.New("paswords does not match")
	}

	return user, err

}
