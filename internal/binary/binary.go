package binary

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

func ToDecimal(binary string) int64 {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Error(err)
		return 0
	}
	return decimal
}
