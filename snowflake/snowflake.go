package snowflake

import (
	"errors"
	"time"

	"github.com/wdy0808/go-common/log"
)

type SnowFlakeGenerator struct {
	dataCenterID   int
	machineID      int
	lastTimeStamp  int64
	sequenceNumber int
}

func NewGenerator(dataCenterNumber, machineNumber int) (out *SnowFlakeGenerator) {
	out = &SnowFlakeGenerator{dataCenterID: dataCenterNumber, machineID: machineNumber}
	return
}

func (this *SnowFlakeGenerator) NextId() (id uint64, err error) {
	now := time.Since(startTimestamp).Nanoseconds()
	if now == this.lastTimeStamp {
		this.sequenceNumber++
		if this.sequenceNumber > maxSequenceNumber {
			log.LogError("snowflake has created [%d] id on timestamp [%d]", maxSequenceNumber, now)
			return 0, errors.New("not id available now")
		}
	} else {
		this.lastTimeStamp = now
		this.sequenceNumber = 0
	}

	return (uint64(this.lastTimeStamp) << timestampShiftLeftBits & timestampBits) |
		(uint64(this.dataCenterID) << dataCenterShiftLeftBits & dataCenterBites) |
		(uint64(this.machineID) << machineShiftLeftBits & machineBites) |
		(uint64(this.sequenceNumber) & sequenceBits), nil
}
