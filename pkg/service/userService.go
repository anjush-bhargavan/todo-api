package service

import (
	"errors"

	"github.com/anjush-bhargavan/todo-api/config"
	"github.com/anjush-bhargavan/todo-api/pkg/models"
	inter "github.com/anjush-bhargavan/todo-api/pkg/repo/interfaces"
	"github.com/anjush-bhargavan/todo-api/pkg/service/interfaces"
	"github.com/anjush-bhargavan/todo-api/utility"
	"github.com/gocql/gocql"
)

type UserService struct {
	Repo inter.UserRepositoryInter
	Cnfg *config.Config
}

func NewUserService(repo inter.UserRepositoryInter, cnfg *config.Config) interfaces.UserServiceInter {
	return &UserService{Repo: repo,
		Cnfg: cnfg}
}

func (s *UserService) UserSignUpSvc(user *models.User) error {
	user.ID = gocql.TimeUUID()
	hashedPassword, err := utility.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.Repo.CreateUser(user)
}

func (s *UserService) UserLoginSvc(login *models.Login) (string, error) {
	user, err := s.Repo.GetUserByEmail(login.Email)
	if err != nil {
		return "", err
	}

	if !utility.CheckPassword(login.Password, user.Password) {
		return "", errors.New("password incorrect")
	}

	token, err := utility.GenerateToken(s.Cnfg.SECRETKEY, user.Email, user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
