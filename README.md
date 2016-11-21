# ipquery

ip query library for golang 

## Usage

```
package main

import (
    "fmt"
    "github.com/tabalt/ipquery"
)

func main() {
  	df := "testdata/test_10000.data"
  	err := ipquery.Load(df)
  	if err != nil {
  		fmt.Println(err)
  	}
    
  	ip := "61.149.208.1"
  	dt, err := ipquery.Find(ip)
  
  	if err != nil {
  		fmt.Println(err)
  	} else {
  		fmt.Println(ip, string(dt))
  	}
}
```

