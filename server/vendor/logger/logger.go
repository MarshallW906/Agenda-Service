package logger

import (
	"log"
	"os"
)

var (
	infoLogger *log.Logger
	errLogger  *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	errLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
}

// Info : Log normal operations (no error occured)
func Info(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

// FatalIf : Log operations with error aroused
func FatalIf(err error) {
	if err != nil {
		errLogger.Fatalln(err)
	}
}

func RecordErr(err error) {
	if err != nil {
		errLogger.Println(err)
	}
}
