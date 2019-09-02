package log

import "log"
import "os"

var _FILE_PATH string
var (
	logInfo    *log.Logger
	logWarning *log.Logger
	logError   *log.Logger
)

func init() {
	err := os.Rename(_FILE_PATH, _FILE_PATH+".backup")
	file, err := os.OpenFile(_FILE_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
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
