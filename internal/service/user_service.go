package service

import (
	"NorthlaneVPN/internal/models"
	"NorthlaneVPN/internal/repository"
	"database/sql"
	"errors"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetUserById(id int) (models.User, error)
	RegisterOrGet(TgId int64) (models.User, error)
	SetBanStatus(status bool, id int) error
}

type UserServiceImpl struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (u *UserServiceImpl) CreateUser(user models.User) (models.User, error) {
	user, err := u.repo.CreateUser(user)
	if err != nil {
		return user, err
	}
	return user, err
}

func (u *UserServiceImpl) GetUserById(id int) (models.User, error) {
	return u.repo.GetUserById(id)
}

func (u *UserServiceImpl) RegisterOrGet(TgId int64) (models.User, error) {
	user, err := u.repo.GetUserByTgId(TgId)

	// Проверяем существует ли telegram id
	if errors.Is(err, sql.ErrNoRows) {
		return u.repo.CreateUser(models.User{TgId: TgId})
	}

	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (u *UserServiceImpl) SetBanStatus(status bool, id int) error {
	return u.repo.SetBanStatus(status, id)
}
