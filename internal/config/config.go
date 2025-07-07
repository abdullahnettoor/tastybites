package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBConfig     DBConfig
	ServerConfig ServerConfig
}

type DBConfig struct {
	Driver string // e.g., "postgres", "mongodb", etc.
	// connection details
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type ServerConfig struct {
	Host string
	Port int
}

func LoadConfig() (*Config, error) {

	dbDriver := os.Getenv("TASTYBITES_DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "postgres" // Default value if not set
	}
	dbHost := os.Getenv("TASTYBITES_DB_HOST")
	if dbHost == "" {
		dbHost = "localhost" // Default value if not set
	}
	dbPortStr := os.Getenv("TASTYBITES_DB_PORT")
	if dbPortStr == "" {
		dbPortStr = "5432" // Default value if not set
	}
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, err // Return error if conversion fails
	}
	dbUsername := os.Getenv("TASTYBITES_DB_USERNAME")
	if dbUsername == "" {
		dbUsername = "postgres" // Default value if not set
	}
	dbPassword := os.Getenv("TASTYBITES_DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres" // Default value if not set
	}
	dbDatabase := os.Getenv("TASTYBITES_DB_DATABASE")
	if dbDatabase == "" {
		dbDatabase = "tastybites" // Default value if not set
	}
	serverHost := os.Getenv("TASTYBITES_SERVER_HOST")
	if serverHost == "" {
		serverHost = "localhost" // Default value if not set
	}
	serverPortStr := os.Getenv("TASTYBITES_SERVER_PORT")
	if serverPortStr == "" {
		serverPortStr = "8080" // Default value if not set
	}
	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil {
		return nil, err // Return error if conversion fails
	}

	// Load configuration from environment variables or a config file
	return &Config{
		DBConfig: DBConfig{
			Driver: dbDriver,
			Host:   dbHost,
			Port:   dbPort,
			Username: dbUsername,
			Password: dbPassword,
			Database: dbDatabase,
		},
		ServerConfig: ServerConfig{
			Host: serverHost,
			Port: serverPort,
		},
	}, nil
}
