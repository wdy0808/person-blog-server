package uuid

import (
	"github.com/google/uuid"
	"github.com/wdy0808/person-blog-server/service/log"
)

func NewRandomUuid() (newUuid string, err error) {
	id, err := uuid.NewRandom()
	time := 0
	for err != nil {
		id, err = uuid.NewRandom()
		time++
		if time >= retryTime {
			log.LogError("create random uuid exceed max retry time [%d], last error message [%s]", retryTime, err.Error())
			panic("random uuid generation error")
		}
	}

	newUuid = id.String()
	return
}
