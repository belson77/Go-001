package config

type Config struct {
	DB  DBConfig
	RPC RPCConfig
}

type DBConfig struct {
	Host     string
	Username string
	Password string
	Dataname string
}

type RPCConfig struct {
	Address string
}
