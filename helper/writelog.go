package helper

import (
	"log"
	"os"
)

func WriteLog(filePath string, errorString string, queryString string) {
	f, err := os.OpenFile("./logs"+filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "logger: ", log.LstdFlags)
	logger.Println(errorString)
	if queryString != "" {
		logger.Println("Query statement: " + queryString)
	}
}
