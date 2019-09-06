package snowflake

import "time"

const (
	timestampBits           uint64 = 0x7fffffffffc00000
	dataCenterBites         uint64 = 0x00000000003e0000
	machineBites            uint64 = 0x000000000001f000
	sequenceBits            uint64 = 0x0000000000000fff
	timestampShiftLeftBits         = 22
	dataCenterShiftLeftBits        = 17
	machineShiftLeftBits           = 12
	maxSequenceNumber              = 4096
)

var startTimestamp = time.Date(2019, 8, 8, 8, 30, 0, 0, time.UTC)
