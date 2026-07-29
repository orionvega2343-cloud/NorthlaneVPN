package repository

import (
	"NorthlaneVPN/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	CreateUser(user models.User) (models.User, error)
	GetUserById(id int) (models.User, error)
	GetUserByTgId(TgId int) (models.User, error)
	SetBanStatus(status bool, id int) error
}

type UserRepoImpl struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

func (r *UserRepoImpl) CreateUser(user models.User) (models.User, error) {
	err := r.db.Get(&user, `INSERT INTO users(username, tg_id, is_banned) VALUES ($1, $2, $3) RETURNING  id, created_at`, user.Username, user.TgId, user.IsBanned)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepoImpl) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, `SELECT id, username, tg_id, created_at, is_banned FROM users WHERE id = $1`, id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepoImpl) GetUserByTgId(TgId int) (models.User, error) {
	var users models.User
	err := r.db.Get(&users, `SELECT id, username, tg_id, created_at, is_banned FROM users WHERE tg_id = $1`, TgId)
	if err != nil {
		return models.User{}, err
	}
	return users, nil
}

func (r *UserRepoImpl) SetBanStatus(status bool, id int) error {
	_, err := r.db.Exec(`UPDATE users SET is_banned=$1 WHERE id = $2`, status, id)
	if err != nil {
		return err
	}
	return nil
}
