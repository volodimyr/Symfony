package app

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pressly/goose/v3"
	"github.com/volodimyr/Symphony/migrations"
)

func NewRedisClient(c RedisStoreConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr + ":" + c.Port,
		Password: c.Pwd,
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("couldn't establish (ping) redis storage due to %v", err)
	}

	return rdb, nil
}

func ConnectPostgresDB(c PostgresDBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", "user="+c.Usr+" password="+c.Pwd+" dbname="+c.Name+" host="+c.Host+" sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed connect to postgres db due to %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("couldn't ping postgres db due to %v", err)
	}

	const maxConns = 10

	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxConns)
	db.SetConnMaxLifetime(1 * time.Minute)

	return db, nil
}

func MigratePostgresDBSchema(db *sql.DB) error {
	goose.SetBaseFS(migrations.Files)
	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("failed to migrate postgres schema due to %v", err)
	}

	return nil
}
