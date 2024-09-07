package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	DB *mongo.Client // instance of mongo db
}

func NewConfig(db *mongo.Client) *Config {
	return &Config{
		DB: db,
	}
}

type ContextKey string

const ConfigKey ContextKey = "config"
