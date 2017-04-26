package context

import (
	"os"
	"strconv"
)

type (
	// Конфигурация
	Configuration struct {
		PsqlConfiguration   *PsqlConfiguration
		ServerConfiguration *ServerConfiguration
	}

	// Конфигурация БД
	PsqlConfiguration struct {
		Host     string
		Port     string
		User     string
		Password string
		DbName   string
		SslMode  string
		IsCreate bool
	}

	// Конфигурация сервера
	ServerConfiguration struct {
		Port     string
		TokenKey string
	}
)

func Build() *Configuration {
	config := &Configuration{
		PsqlConfiguration:   buildPsqlConfiguration(),
		ServerConfiguration: buildServerConfiguration(),
	}
	return config
}

func buildPsqlConfiguration() *PsqlConfiguration {
	isCreate, _ := strconv.ParseBool(lookupEnvOrDefault("DB_INITIALIZATION", "false"))
	return &PsqlConfiguration{
		Host:     lookupEnvOrDefault("DB_HOST", "127.0.0.1"),
		Port:     lookupEnvOrDefault("DB_PORT", "5432"),
		DbName:   lookupEnvOrDefault("DB_DATABASE", ""),
		User:     lookupEnvOrDefault("DB_USER", "postgres"),
		Password: lookupEnvOrDefault("DB_PASSWORD", "postgres"),
		SslMode:  lookupEnvOrDefault("DB_SSL", "disable"),
		IsCreate: isCreate,
	}
}

func buildServerConfiguration() *ServerConfiguration {
	return &ServerConfiguration{
		Port: lookupEnvOrDefault("SERVER_PORT", ":8080"),
		TokenKey: lookupEnvOrDefault("TOKEN_KEY", ""),
	}
}

func lookupEnvOrDefault(key string, other string) string {
	value, present := os.LookupEnv(key)
	if !present {
		value = other
	}
	return value
}
