package config

import (
	"encoding/json"
	"github.com/chrisheath/time-tracker/pkg/logger"
	"go/build"
	"io/ioutil"
	"os"
)

type Config interface {
	Load()
}

type AppConfig struct {
	Logs, FileFormat, LogFormat string
}

func GetConfig() AppConfig {
	return AppConfig{}
}

func (a *AppConfig) Load() {
	jsonFile, err := os.Open(build.Default.GOPATH + "/src/github.com/chrisheath/time-tracker/config.json")
	if err != nil {
		logger.Write([]string{"ERROR:", string(err.Error())})
	}
	defer jsonFile.Close()

	appConfigData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logger.Write([]string{"ERROR:", string(err.Error())})
	}

	err = json.Unmarshal(appConfigData, a)
	if err != nil {
		logger.Write([]string{"ERROR:", string(err.Error())})
	}

	logger.Write([]string{"INFO:", "LOADED CONFIG :-"})
	logger.Write([]string{"INFO:", "logs: ", a.Logs})
	logger.Write([]string{"INFO:", "fileFormat: ", a.FileFormat})
	logger.Write([]string{"INFO:", "logFormat: ", a.LogFormat})
}
