package simplehttp

import (
	"fmt"
	"testing"
)

func TestHTTP(t *testing.T) {
	jsresp, err := GetResponeBody("http://httpbin.org/ip")
	if err != nil {
		t.FailNow()
	}
	fmt.Printf("IP-Json:\n%v", jsresp)
}
