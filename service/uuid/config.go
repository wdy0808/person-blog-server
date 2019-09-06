package uuid

var retryTime = 10

func GetRetryTime() int {
	return retryTime
}

func SetRetryTime(time int) {
	retryTime = time
}
