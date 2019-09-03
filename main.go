package main
 
import "os"
import "os/signal"
import "syscall"
 
func main() {
	signalChan := make(chan os.Signal,1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
