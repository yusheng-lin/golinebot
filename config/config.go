package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	LineChannelSecret string `mapstructure:"LINE_CHANNEL_SECRET"`
	LineChannelToken  string `mapstructure:"LINE_CHANNEL_ACCESS_TOKEN"`
	MongoURL          string `mapstructure:"MONGODB_URL"`
	MongoDB           string `mapstructure:"MONGODB_NAME"`
	MongoUser         string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
	MongoPwd          string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`
}

func NewConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	files, err := os.ReadDir(path)

	if err != nil {
		return
	}
	config = &Config{}
	for _, file := range files {
		viper.SetConfigFile(fmt.Sprintf("%s/%s", path, file.Name()))
		err = viper.ReadInConfig()
		if err != nil {
			return
		}
		err = viper.Unmarshal(config)
	}
	return
}
