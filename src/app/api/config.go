package api

import "os"

type Config struct {
	DB *DBConfig
}
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
	PortNo   int
	Host     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: os.Getenv("DATABASE_USERNAME"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_SCHEMA"),
			Host:    os.Getenv("DATABASE_HOST"),
			Charset:  "utf8",
			PortNo:   5432,
		},
	}
}
