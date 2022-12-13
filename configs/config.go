package configs

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Environment            string
	DynamoTransactionTable string
	DynamoEndpoint         string
	DynamoRegion           string
}

func GetAppConfigs() AppConfig {
	env := os.Getenv("environment")
	if env == "" {
		log.Fatalf("Failed to read environment config [%s]", env)
	}

	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigType("yaml")
	v.SetConfigName(env)
	v.AddConfigPath("../../configs")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file, error: %s", err.Error())
	}

	appConfig := extractConfigVars(v)
	appConfig.Environment = env

	return appConfig
}

func extractConfigVars(v *viper.Viper) AppConfig {
	appConfig := AppConfig{}

	appConfig.DynamoTransactionTable = v.GetString("dynamo.transactionTable")
	appConfig.DynamoEndpoint = v.GetString("dynamo.endpoint")
	appConfig.DynamoRegion = v.GetString("dynamo.region")

	return appConfig
}
