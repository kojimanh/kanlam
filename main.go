package main

import (
	"kanlam/lib/config"
	"kanlam/server"
	storage "kanlam/storage/maindb"
)

func main() {
	config.SetupConfig()
	storage.ConnectMainDb()
	server.StartServer()

	defer storage.MainPool.Close()
}
