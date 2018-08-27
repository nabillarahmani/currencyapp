package configs

type (
	// Configuration struct holds all configuration from config.yml
	Configuration struct {
		Server   ServerConfiguration
		Database DatabaseConfiguration
	}

	// ServerConfiguration holds server configurations
	ServerConfiguration struct {
		Port string
	}

	// DatabaseConfiguration holds database configurations
	DatabaseConfiguration struct {
		ConnectionURI string
	}
)
