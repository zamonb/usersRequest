package logger

import (
	"log"
	"os"
)

var(
	Info *log.Logger
)

func Init() {
	logFileInfo := "logfile.log"

	fileInfo, err := os.OpenFile(logFileInfo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Info = log.New(fileInfo, "", log.Ldate|log.Lmicroseconds)
	Info.SetOutput(fileInfo)
	Info.Println("service launched")
}
