package config

import (
	"log"

	"github.com/spf13/viper"
)

func bindEnv(confName, envName string, defaultVal any) {
	viper.SetDefault(confName, defaultVal)
	if err := viper.BindEnv(confName, envName); err != nil {
		log.Fatal(err)
	}
}

func init() {
	bindEnv("OpensearchURL", "OPENSEARCH_URL", "http://localhost:9200")
	bindEnv("ServerURL", "SERVER_URL", "localhost:8080")
}

func OpensearchURL() string {
	return viper.GetString("OpensearchURL")
}

func ServerURL() string {
	return viper.GetString("ServerURL")
}
