package userstorage

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	storage "kanlam/storage/maindb"
)

func (r StructUserStorage) GetByUsername(username string) (*User, error) {
	sql := "SELECT id, username, password FROM public.user WHERE username=$1"

	user := new(User)

	err := storage.MainPool.QueryRow(context.Background(), sql, username).Scan(&user.Id, &user.Username, &user.Password)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
