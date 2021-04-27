package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var configFilename = "../.env"

func init() {
	viper.SetConfigFile(configFilename)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Connont find config file, %s", err)
	}
}

func NewConfig() (defConfig *Config, err error) {
	defConfig = &Config{}

	appEnv := viper.GetString("APP_ENV")
	appPort := viper.GetInt("APP_PORT")
	appDebug := viper.GetBool("APP_DEBUG")

	dbDriverName := viper.GetString("DB_DRIVER")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetInt("DB_PORT")
	dbName := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")

	if appEnv == "" || appPort == 0 {
		err = fmt.Errorf("[CONFIG][Critical] Please check section APP on %s", configFilename)
		return
	}

	defConfig.AppEnv = appEnv
	defConfig.AppPort = appPort
	defConfig.Debug = appDebug

	if dbHost == "" || dbPort == 0 || dbUser == "" || dbName == "" || dbDriverName == "" {
		err = fmt.Errorf("[CONFIG][Critical] Please check section DB on %s", configFilename)
		return
	}

	dbConfig := &DB{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUser,
		Password: dbPassword,
		Name:     dbName,
		Driver:   dbDriverName,
	}

	defConfig.DB = dbConfig

	return defConfig, nil
}
