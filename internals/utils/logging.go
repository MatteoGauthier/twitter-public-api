package logging

import (
	"log"
	"os"
)

var info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var warn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
var error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

func InfoLog(message string) {
	info.Println(message)
}

func WarningLog(message string) {
	warn.Println(message)
}

func ErrorLog(message string) {
	error.Println(message)
}
