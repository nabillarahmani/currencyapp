package configs

import (
	// library used from modules
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
