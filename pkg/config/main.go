package config

import (
	"encoding/json"
	"github.com/chrisheath/time-tracker/pkg/logger"
	"go/build"
	"io/ioutil"
	"os"
)

type AppConfig struct {
	Logs, FileFormat, LogFormat string
}

func Load() AppConfig {
	appConfig := AppConfig{}
	configPath := build.Default.GOPATH + "/src/github.com/chrisheath/time-tracker/config.json"

	jsonFile, err := os.Open(configPath)
	if err != nil {
		logger.Write([]string{"ERROR:", string(err.Error())})
	}
	defer jsonFile.Close()

	appConfigData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logger.Write([]string{"ERROR:", string(err.Error())})
	}

	err = json.Unmarshal(appConfigData, appConfig)
	if err != nil {
		logger.Write([]string{"ERROR:", string(err.Error())})
	}

	logger.Write([]string{"INFO:", "LOADED CONFIG :-"})
	logger.Write([]string{"INFO:", "logs: ", appConfig.Logs})
	logger.Write([]string{"INFO:", "fileFormat: ", appConfig.FileFormat})
	logger.Write([]string{"INFO:", "logFormat: ", appConfig.LogFormat})

	return appConfig
}
