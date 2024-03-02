package storageping

import (
	"context"
	storage "kanlam/storage/maindb"
)

func PingDb() (string, error) {
	var result string
	err := storage.MainPool.QueryRow(context.Background(), "select 1 +1 as result").Scan(&result)

	return result, err
}
