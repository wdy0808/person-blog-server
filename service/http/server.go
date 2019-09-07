package http

import (
	"time"

	"github.com/wdy0808/go-common/log"
	"github.com/wdy0808/httpgo"
	"github.com/wdy0808/person-blog-server/service/conf"
)

type serverConfig struct {
	port          int64
	maxReadTime   time.Duration
	maxWriteTime  time.Duration
	allowedOrigin []string
}

var config serverConfig

func init() {
	configInfomation := conf.GetConfigInformation("server")
	config.port = configInfomation.MustInt64("port")
	config.allowedOrigin = configInfomation.MustStringArray("allow_origin_array")
	timeDuration, err := time.ParseDuration(configInfomation.MustString("write_time_out"))
	if err != nil {
		log.LogError("parse time duration [%s] error [%s]", configInfomation.MustString("write_time_out"), err.Error())
		return
	}
	config.maxWriteTime = timeDuration
	timeDuration, err = time.ParseDuration(configInfomation.MustString("read_time_out"))
	if err != nil {
		log.LogError("parse time duration [%s] error [%s]", configInfomation.MustString("read_time_out"), err.Error())
		return
	}
	config.maxReadTime = timeDuration
}

func LitenAndServe() {
	httpgo.RunForver(httpgo.ServerConfig{
		Port:          config.port,
		ReadTimeout:   config.maxReadTime,
		WriteTimeout:  config.maxWriteTime,
		AllowedOrigin: config.allowedOrigin,
	})
}
