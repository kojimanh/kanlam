package taskstorage

import (
	"context"
	storage "kanlam/storage/maindb"
)

func (r StructTaskStorage) CreateTask(task Task) error {
	sql := "INSERT INTO public.task (id, user_id, name, content) VALUES ($1, $2, $3)"

	_, err := storage.MainPool.Exec(context.Background(), sql, task.Id, task.UserId, task.Name, task.Content)

	return err
}
