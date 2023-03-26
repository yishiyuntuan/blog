package config

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func NewConfig() *Config {
	var configFile string
	flag.StringVar(&configFile, "config", "", "")
	flag.Parse()

	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType("yaml")
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("./conf/")
		viper.SetConfigName("config")
	}

	conf := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}
	return conf
}

func initDirectory(conf *Config) {
	mkdirFunc := func(dir string, err error) error {
		if err == nil {
			if _, err = os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, os.ModePerm)
			}
		}
		return err
	}
	err := mkdirFunc(conf.Sonic.LogDir, nil)
	err = mkdirFunc(conf.Sonic.UploadDir, err)
	if err != nil {
		panic(fmt.Errorf("initDirectory err=%w", err))
	}
}

var mode string

func IsDev() bool {
	return mode == "development"
}
