package ipquery

import (
	"errors"
)

const (
	ipRangeFieldCount = 3
)

var ErrorIpRangeNotFound = errors.New("ip range not found")

type IpRange struct {
	Begin uint32
	End   uint32
	Data  []byte
}
