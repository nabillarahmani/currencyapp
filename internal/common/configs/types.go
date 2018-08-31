package configs

import (
	"github.com/nabillarahmani/currencyapp/internal/common/database"
)

type (
	// Configuration struct holds all configuration from config.yml
	Configuration struct {
		// Configs
		Server   ServerConfiguration
		Database DatabaseConfiguration

		// Object conns
		db DatabaseObjects
	}

	// ServerConfiguration holds server configurations
	ServerConfiguration struct {
		Port string
	}

	// DatabaseConfiguration holds database configurations
	DatabaseConfiguration struct {
		ConnectionURI  string
		ConnectionType string
	}

	// DatabaseObjects is
	DatabaseObjects struct {
		Conns database.Database
	}
)
