package log

import (
	"log"
	"os"

	"github.com/wdy0808/person-blog-server/service/file"
)

var (
	logInfo    *log.Logger
	logWarning *log.Logger
	logError   *log.Logger
)

func init() {
	currentDir := file.CurrentDir()
	logFilePath := currentDir + "/log/logfile"
	if file.FileExist(logFilePath) {
		os.Rename(logFilePath, logFilePath+".backup")
	}
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err.Error())
	}

	logInfo = log.New(file, "[INFO] ", log.Ldate|log.Lmicroseconds)
	logWarning = log.New(file, "[WARNING] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	logError = log.New(file, "[ERROR] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
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
