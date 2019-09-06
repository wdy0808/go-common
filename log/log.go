package log

import (
	"log"
	"os"
)

var (
	logInfo    *log.Logger
	logWarning *log.Logger
	logError   *log.Logger
)

func init() {
	logInfo = log.New(os.Stderr, "[INFO] ", log.Ldate|log.Lmicroseconds)
	logWarning = log.New(os.Stderr, "[WARNING] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	logError = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

func LogInfo(format string, a ...interface{}) {
	logInfo.Printf(format, a...)
}

func LogWarning(format string, a ...interface{}) {
	logWarning.Printf(format, a...)
}

func LogError(format string, a ...interface{}) {
	logError.Printf(format, a...)
}
