package config

import (
	"GoDocRestPostgres/api"
	"database/sql"
	"log"
)

const (
	DB_USER     = "docker"
	DB_PASSWORD = "docker"
	DB_NAME     = "docker"
	PORT        = "5432"
)

//Configuration structs holds all our components of the app
type Configuration struct {
	Client api.RecipeClient
	Logger *log.Logger
}

// NewConfiguration injects the config struct with our instances of logger, db etc
func NewConfiguration(db *sql.DB, log *log.Logger) *Configuration {

	client := &api.RecipeImpl{Db: db}

	Conf := Configuration{
		Client: client,
		Logger: log,
	}

	return &Conf
}
