package main

import (
	"encoding/json"
	//"io"
	"io/ioutil"
	"os"
	//"path/filepath"
	"github.com/chrisheath/time-tracker/pkg/logger"
)

type AppConfig struct {
	Logs, FileFormat, LogFormat string
}

func main() {
	loadConfigFile()

	// check if file exists for today
	// create file for today with timestamp
	// check todays file for last action (last at bottom?)
	// set current status to be last action
	// add new entry with change in action/project and timestamp

	// content := "test content"
	// // create file to write to
	// file, err := os.Create("./test.txt")
	// checkError(err)
	// // close always close
	// defer file.Close()
	// // write to file
	// ln, err := io.WriteString(file, content)
	// checkError(err)
	// fmt.Println("finished", ln)
	// // read file content
	// newContent, err := ioutil.ReadFile("./test.txt")
	// checkError(err)
	// fmt.Println("from file: ", newContent)
	// result := string(content)
	// fmt.Println("converted from bytes: ", result)
	// // time string
	// t := time.Now()
	// t.String()
	// timeNow := t.Format("2006-01-02--15-04-05")
	// fmt.Println("timeNow: ", timeNow)
	// // file path
	// root, _ := filepath.Abs(".")
	// fmt.Println("\nProcessing path: ", root)
	// fileErr := filepath.Walk(root, processPath)
	// checkError(fileErr)
	// decoding json string
	// decoder := json.NewDecoder(strings.NewReader(content))
	// _, err := decoder.Token()
	// checkError(err)

	defer logger.Write([]string{"INFO:", "Finished!"})
}

// func processPath(path string, info os.FileInfo, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	if path != "." {
// 		if info.IsDir() {
// 			fmt.Println("Directory: ", path)
// 		} else {
// 			fmt.Println("File:", path)
// 		}
// 	}
// 	return nil
// }

func loadConfigFile() {
	jsonFile, err := os.Open("./config.json")
	if err != nil {
		logger.Write([]string{"ERROR:", string(err.Error())})
	}
	defer jsonFile.Close()

	appConfigData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logger.Write([]string{"ERROR:", string(err.Error())})
	}

	var appConfig AppConfig

	err = json.Unmarshal(appConfigData, &appConfig)
	if err != nil {
		logger.Write([]string{"ERROR:", string(err.Error())})
	}

	logger.Write([]string{"INFO:", "LOADED CONFIG :-"})
	logger.Write([]string{"INFO:", "logs: ", appConfig.Logs})
	logger.Write([]string{"INFO:", "fileFormat: ", appConfig.FileFormat})
	logger.Write([]string{"INFO:", "logFormat: ", appConfig.LogFormat})
}
