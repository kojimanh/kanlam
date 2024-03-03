package userstorage

import (
	"context"
	"errors"
	storage "kanlam/storage/maindb"
)

func (r StructUserStorage) GetByUsername(username string) (*User, error) {
	sql := "SELECT id, username, password FROM public.user WHERE username=$1"

	user := new(User)

	if err := storage.MainPool.QueryRow(context.Background(), sql, username).Scan(&user.Id, &user.Username, &user.Password); err != nil {
		return user, errors.New("User not existed")
	}

	return user, nil
}
