package uuid

import "github.com/google/uuid"

func NewSpaceDataUuid(spaceName, data string) (newUuid string) {
	var dataByte []byte = []byte(data)
	space := uuid.Must(uuid.Parse(spaceName))
	newUuid = uuid.NewSHA1(space, dataByte).String()
	return
}