package uuid

import "github.com/google/uuid"

func NewRandomUuid() (newUuid string, err error) {
	id, err := uuid.NewRandom()
	time := 0
	for err != nil {
		id, err = uuid.NewRandom()
		time++
		if time >= retryTime {
			return "", err
		}
	}

	newUuid = id.String()
	return;
}