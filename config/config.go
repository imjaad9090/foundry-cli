package config

import (
	"os"

	"foundry/cli/logger"

	"github.com/spf13/viper"
)

func Init() error {
	// /Users/vasekmlejnsky/Library/Application Support/foundrycli
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}


	// configDir = "/Users/vasekmlejnsky/Developer"

	dirPath := configDir + "/foundrycli"
	confName := "config"
	ext := "json"
	fullPath := dirPath + "/" + confName + "." + ext

	viper.SetConfigName(confName)
	viper.SetConfigType(ext)
	viper.AddConfigPath(dirPath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm)

		f, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		defer f.Close()
		f.WriteString("{}")
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}

	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func Set(key string, val interface{}) {
	logger.Fdebugln("Set to config (key, val)", key, val)
	viper.Set(key, val)
}

func GetString(key string) string {
	logger.Fdebugln("Get string from config:", key)
	return viper.GetString(key)
}

func GetInt(key string) int {
	logger.Fdebugln("Get int from config:", key)
	return viper.GetInt(key)
}

func Write() error {
	logger.Fdebugln("Write config")
	return viper.WriteConfig()
}
