package userstorage

import (
	"context"
	storage "kanlam/storage/maindb"
)

func (r StructUserStorage) CreateUser(user User) error {
	sql := "INSERT INTO public.user VALUES ($1, $2, $3)"

	_, err := storage.MainPool.Exec(context.Background(), sql, user.Id, user.Username, user.Password)

	return err
}
