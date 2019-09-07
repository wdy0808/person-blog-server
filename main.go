package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/wdy0808/person-blog-server/service/http"
)

func main() {
	http.LitenAndServe()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
