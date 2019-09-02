package jason

import "os"
import "strings"
import "path/filepath"
import jason "gopkg.in/antonholmquist/jason.v1"
import "github.com/wdy0808/person-blog-server/service/log"

type ConfigInformation struct {
	object *jason.Object
}

var _FILE_PATH string
var configMap map[string]*ConfigInformation

const configFilePatern = "*.conf.json"

func init() {
	configMap = make(map[string]*ConfigInformation)
	files, err := filepath.Glob(_FILE_PATH + configFilePatern)

	if nil != err {
		log.LogError("try to get config file error [%s]\n", err.Error())
		panic("jason error")
	}

	if nil == files {
		return
	}

	for _, file := range files {
		fileReader, err := os.Open(file)
		if nil != err {
			log.LogError("read config file [%s] error [%s]\n", file, err.Error())
			panic("jason error")
		}

		object, err := jason.NewObjectFromReader(fileReader)
		if nil != err {
			log.LogError("jason parse config file [%s] error [%s]\n", file, err.Error())
			panic("jason error")
		}

		configType := strings.TrimSuffix(file, strings.TrimPrefix(configFilePatern, "*"))
		configMap[configType] = &ConfigInformation{object}
	}
	return
}

func GetConfigInformation(namespace string) (out *ConfigInformation) {
	out, ifExist := configMap[namespace]
	if !ifExist {
		log.LogError("get config [%s] error [%s]\n", namespace, err.Error())
		panic("jason error")
	}
	return 
}

func (this *ConfigInformation) MustObject(key string) *ConfigInformation {
	valueObject, err := this.object.GetObject(key)
	this.object = valueObject
	if nil != err {
		log.LogError("config file get value [%s] error [%s]\n", key, err.Error())
		panic("jason error")
	}
	return this
}

func (this *ConfigInformation) MustObjectArray(key string) (out []*ConfigInformation) {
	valueObjects, err := this.object.GetObjectArray(key)
	if nil != err {
		log.LogError("config file get value [%s] error [%s]\n", key, err.Error())
		panic("jason error")
	}
	out = make([]*ConfigInformation, len(valueObjects))
	for _, valueObject := range valueObjects {
		config := &ConfigInformation{valueObject}
		out = append(out, config)
	}
	return
}

func (this *ConfigInformation) MustString(key string) (out string) {
	out, err := this.object.GetString(key)
	if nil != err {
		log.LogError("config file get value [%s] error [%s]\n", key, err.Error())
		panic("jason error")
	}
	return
}

func (this *ConfigInformation) MustInt64(key string) (out int64) {
	out, err := this.object.GetInt64(key)
	if nil != err {
		log.LogError("config file get value [%s] error [%s]\n", key, err.Error())
		panic("jason error")
	}
	return
}
