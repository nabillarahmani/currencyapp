package configs

import (
	// library used from modules
	"github.com/nabillarahmani/currencyapp/internal/common/database"
	"github.com/nabillarahmani/currencyapp/internal/common/log"

	// library used from outside
	"github.com/spf13/viper"
)

// GlobalConfig is global config to init many configurations
var GlobalConfig Configuration

// InitConfig is a function to scan variables to struct
func InitConfig() {
	viper.SetConfigName("config")              // name of config file (without extension)
	viper.AddConfigPath("./../files/configs/") // path to look for the config file in
	viper.AddConfigPath(".")                   // optionally look for config in the working directory
	err := viper.ReadInConfig()                // Find and read the config file
	if err != nil {                            // Handle errors reading the config file
		log.Fatal("Error while reading config from config.yml")
		return
	}

	// unmarshall the configs to
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		log.Fatal("unable to decode configs var into struct")
		return
	}
	return
}

// InitDatabaseConn is a function to init and store db obj to global var
func InitDatabaseConn() {
	if GlobalConfig.Database.ConnectionURI == "" {
		log.Fatal("No valid db conn host exist!")
	}

	GlobalConfig.db.Conns = database.Init(GlobalConfig.Database.ConnectionType, GlobalConfig.Database.ConnectionURI)
}

// GetDatabaseObj will return database obj
func GetDatabaseObj() (db database.Database) {
	db = GlobalConfig.db.Conns
	return
}
