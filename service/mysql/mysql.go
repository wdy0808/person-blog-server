package mysql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/wdy0808/person-blog-server/service/jason"
import "github.com/wdy0808/person-blog-server/service/log"

func init() {
	db, err := sql.Open("mysql", jason.GetConfigInformation("mysql").MustString("source"))
	if nil != err {
		log.LogError("open mysql source error [%s]", err.Error())
	}
	err = db.Ping()
	if nil != err {
		log.LogError("ping mysql error [%s]", err.Error())
	}

	stat := db.Stats()
	log.LogInfo("mysql111 connection [%d]", stat.OpenConnections)
}
