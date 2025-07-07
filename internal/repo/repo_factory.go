package repo

import (
	"errors"

	"github.com/abdullahnettoor/tastybites/internal/config"
	"github.com/abdullahnettoor/tastybites/internal/repo/interfaces"
	pgrepo "github.com/abdullahnettoor/tastybites/internal/repo/postgres"
)

func NewRepository(dbConfig *config.DBConfig) (interfaces.Repository, error) {

	switch dbConfig.Driver {
	case "postgres":
		return pgrepo.NewRepository(dbConfig)
	default:
		return nil, errors.New("unsupported database driver: " + dbConfig.Driver)
	}
}
