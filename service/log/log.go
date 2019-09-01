package log

import "log"
import "os"

var (
	LogInfo    *log.Logger
	LogWarning *log.Logger
	LogError   *log.Logger
)

func init()  {
	file, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
    	log.Fatalf("file open error : %v", err)
	}
	
	LogInfo =    log.New(file, "[INFO] ", log.Ldate | log.Lmicroseconds)
	LogWarning = log.New(file, "[WARNING] ", log.Ldate | log.Lmicroseconds | log.Lshortfile)
	LogError =   log.New(file, "[ERROR] ", log.Ldate | log.Lmicroseconds | log.Lshortfile)
}