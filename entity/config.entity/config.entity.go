package entity

type ServerConfig struct {
	Server string
}

type DbConfig struct {
	Url string
}

type ConfigKeyMap struct {
	ServerHost  string
	DbConfigUrl string
}
