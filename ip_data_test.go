package ipquery

import (
	"sync"
	"testing"
)

func TestIpData_Load(t *testing.T) {
	ipData := NewIpData()

	// test load data file
	df := "testdata/test_10.data"
	length := 10

	err := ipData.Load(df)
	if err != nil {
		t.Fatalf("load exists data file failed: %s.", err)
	}

	if ipData.Length() != length {
		t.Fatalf("ip data length error. expected: %d, got: %d.", length, ipData.Length())
	}

	// test reload data file
	df = "testdata/test_10000.data"
	length = 10000

	err = ipData.ReLoad(df)
	if err != nil {
		t.Fatalf("reload data file failed: %s.", err)
	}

	if ipData.Length() != length {
		t.Fatalf("ip data length error. expected: %d, got: %d.", length, ipData.Length())
	}

	// test load not exists data file
	df = "testdata/not_exists.data"
	err = ipData.Load(df)
	if err == nil {
		t.Fatalf("load not exists data file must be failed.")
	}
}

var findCases = []struct {
	ip   string
	err  error
	data []byte
}{
	{"127.0.0.1", ErrorIpRangeNotFound, []byte("")},
	{"192.168.1.1", ErrorIpRangeNotFound, []byte("")},
	{"10.16.10.10", ErrorIpRangeNotFound, []byte("")},
	{"113.68.244.209", ErrorIpRangeNotFound, []byte("")},
	{"121.27.251.247", ErrorIpRangeNotFound, []byte("")},
	{"218.30.116.8", ErrorIpRangeNotFound, []byte("")},
	{"218.30.116", ErrorIpRangeNotFound, []byte("")},
	{"-", ErrorIpRangeNotFound, []byte("")},
	{"153.19.50.62", ErrorIpRangeNotFound, []byte("")},
	{"61.149.208.1", nil, []byte("北京\t北京\t海淀\t101010200\t联通")},
	{"1.25.47.240", nil, []byte("内蒙古\t乌兰察布\t集宁\t101080401\t联通")},
	{"1.85.159.255", nil, []byte("陕西\t西安\t西安\t101110101\t电信")},
}

func TestIpData_Find(t *testing.T) {
	ipData := NewIpData()
	df := "testdata/test_10000.data"
	err := ipData.Load(df)
	if err != nil {
		t.Fatalf("load exists data file failed: %s.", err)
	}

	for _, c := range findCases {
		dt, err := ipData.Find(c.ip)
		if err != c.err {
			t.Errorf("find data for %s expected error: %v, got: %v.", c.ip, c.err, err)
		}

		if err == nil && string(dt) != string(c.data) {
			t.Errorf("find data for %s failed. expected: %s, got: %s.", c.ip, string(c.data), string(dt))
		}
	}
}

func TestIpData_Parallel_Find(t *testing.T) {
	t.Parallel()
	var wg sync.WaitGroup

	ipData := NewIpData()
	df := "testdata/test_10000.data"
	err := ipData.Load(df)
	if err != nil {
		t.Fatalf("load exists data file failed: %s.", err)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for _, c := range findCases {
				dt, err := ipData.Find(c.ip)
				if err != c.err {
					t.Errorf("find data for %s expected error: %v, got: %v.", c.ip, c.err, err)
				}

				if err == nil && string(dt) != string(c.data) {
					t.Errorf("find data for %s failed. expected: %s, got: %s.", c.ip, string(c.data), string(dt))
				}
			}
		}()
	}

	wg.Wait()
}

func TestIpData_ip2Long(t *testing.T) {
	ipData := NewIpData()
	var cases = []struct {
		ip   string
		long uint32
	}{
		{"127.0.0.1", 2130706433},
		{"192.168.1.1", 3232235777},
		{"10.16.10.10", 168823306},
		{"113.68.244.209", 1900344529},
		{"121.27.251.247", 2031877111},
		{"218.30.116.8", 3659428872},
		{"218.30.116", 0},
		{"-", 0},
	}

	for _, c := range cases {
		il := ipData.ip2Long(c.ip)
		if il != c.long {
			t.Errorf("ip2long for %s failed. expected %d, got %d.", c.ip, c.long, il)
		}
	}
}

func BenchmarkIpData_Load(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ipData := NewIpData()
		err := ipData.Load("testdata/ip_chunzhen.txt")
		if err != nil {
			b.Fatalf("load exists data file failed: %s.", err)
		}
	}
}

func BenchmarkIpData_Find(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	ipData := NewIpData()
	err := ipData.Load("testdata/ip_chunzhen.txt")
	if err != nil {
		b.Fatalf("load exists data file failed: %s.", err)
	}

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		ipData.Find("116.5.166.98")
		b.StopTimer()
	}
}
