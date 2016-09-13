package ipquery

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"net"
	"os"
	"strconv"
	"strings"
)

type IpData []*IpRange

func NewIpData() *IpData {
	return &IpData{}
}

//TODO 初始化后对数据做排序
func (id *IpData) Load(df string) error {
	file, err := os.Open(df)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		item := strings.SplitN(line, "\t", ipRangeFieldCount)
		if len(item) != ipRangeFieldCount {
			continue
		}

		begin, _ := strconv.Atoi(item[0])
		end, _ := strconv.Atoi(item[1])
		if begin > end {
			continue
		}

		ir := &IpRange{
			Begin: uint32(begin),
			End:   uint32(end),
			Data:  []byte(item[2]),
		}

		*id = append(*id, ir)
	}

	return scanner.Err()
}

func (id *IpData) ReLoad(df string) error {
	nid := NewIpData()
	err := nid.Load(df)
	if err != nil {
		return err
	}

	*id = *nid
	return nil
}

func (id *IpData) Length() int {
	return len(*id)
}

func (id *IpData) Find(ip string) ([]byte, error) {
	ir, err := id.getIpRange(ip)
	if err != nil {
		return nil, err
	}

	return ir.Data, nil
}

func (id *IpData) getIpRange(ip string) (*IpRange, error) {
	var low, high int = 0, (id.Length() - 1)

	ipdt := *id
	il := ip2Long(ip)
	if il <= 0 {
		return nil, ErrorIpRangeNotFound
	}

	for low <= high {
		var middle int = (high-low)/2 + low

		ir := ipdt[middle]

		if il >= ir.Begin && il <= ir.End {
			return ir, nil
		} else if il < ir.Begin {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return nil, ErrorIpRangeNotFound
}

func ip2Long(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}
