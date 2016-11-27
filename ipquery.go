package ipquery

import (
	"os"
)

var defaultIpData *IpData

func init() {
	defaultIpData = NewIpData()
}

func Load(df string) error {
	reader, err := os.Open(df)
	if err != nil {
		return err
	}
	return defaultIpData.Load(reader)
}

func ReLoad(df string) error {
	reader, err := os.Open(df)
	if err != nil {
		return err
	}
	return defaultIpData.ReLoad(reader)
}

func Length() int {
	return defaultIpData.Length()
}

func Find(ip string) ([]byte, error) {
	ir, err := defaultIpData.Find(ip)
	if err != nil {
		return nil, err
	} else {
		return ir.Data, nil
	}
}
