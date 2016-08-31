package ipquery

import ()

var defaultIpData *IpData

func init() {
	defaultIpData = NewIpData()
}

func Load(df string) error {
	return defaultIpData.Load(df)
}

func ReLoad(df string) error {
	return defaultIpData.ReLoad(df)
}

func Length() int {
	return defaultIpData.Length()
}

func Find(ip string) ([]byte, error) {
	return defaultIpData.Find(ip)
}
