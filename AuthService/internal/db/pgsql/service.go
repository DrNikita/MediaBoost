package pgsql

import (
	"auth-service/internal/db/model"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

func NewPgService(db *sql.DB) UserStore {
	return UserStore{
		db: db,
	}
}

func (us *UserStore) CreateUser(user *model.User) (int64, error) {
	return 0, nil
}

func (us *UserStore) UpdateUser(user *model.User) (int64, error) {
	return 0, nil
}

func (us *UserStore) DeleteUser(id int64) error {
	return nil
}

func (us *UserStore) FindUserByEmailOrNickname(identificator string) (*model.User, error) {
	return nil, nil
}