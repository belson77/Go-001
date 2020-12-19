package config

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string
	Username string
	Password string
	Dataname string
}
