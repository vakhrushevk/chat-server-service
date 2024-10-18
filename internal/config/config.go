package config

import "github.com/joho/godotenv"

type PgConfig interface {
	DSN() string
}

type GRPCConfig interface {
	Address() string
}

// Load - Load config from path
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}
