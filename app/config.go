package app

import "os"

type (
	Config struct {
		Server   HTTPServerConfig
		Postgres PostgresDBConfig
		Redis    RedisStoreConfig
	}

	HTTPServerConfig struct {
		Port string
	}

	PostgresDBConfig struct {
		Host string
		Port string
		Name string
		Usr  string
		Pwd  string
	}

	RedisStoreConfig struct {
		Addr string
		Port string
		Pwd  string
	}
)

func NewConfig() Config {
	return Config{
		Server: HTTPServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},

		Postgres: PostgresDBConfig{
			Host: os.Getenv("POSTGRES_HOST"),
			Port: os.Getenv("POSTGRES_PORT"),
			Name: os.Getenv("POSTGRES_DB"),
			Usr:  os.Getenv("POSTGRES_USER"),
			Pwd:  os.Getenv("POSTGRES_PASSWORD"),
		},

		Redis: RedisStoreConfig{
			Addr: os.Getenv("REDIS_ADDR"),
			Port: os.Getenv("REDIS_PORT"),
			Pwd:  os.Getenv("REDIS_PWD"),
		},
	}
}
