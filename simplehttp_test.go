package simplehttp

import (
	"fmt"
	"testing"
)

func TestHTTP(t *testing.T) {
	jsresp, err := GetResponseBody("http://httpbin.org/ip")
	if err != nil {
		fmt.Print(err)
		t.FailNow()
	}
	fmt.Printf("IP-Json:\n%v", string(jsresp))
}

func TestExternalIP(t *testing.T) {
	ext, err := GetExternalIP()
	if err != nil {
		fmt.Print(err)
		t.FailNow()
	}
	fmt.Printf("IP-Adress: %v\n", ext)
}

func TestWAI(t *testing.T) {
	hn, err := GetWhoAmI()
	if err != nil {
		fmt.Print(err)
		t.FailNow()
	}
	fmt.Printf("I am: %v\n", hn)
}

func TestIdentity(t *testing.T) {
	var id Identity
	id = Identity{"localhost", "127.0.0.1"}
	if "simplehttp.Identity: localhost (127.0.0.1)" != fmt.Sprintf("%T: %v", id, id) {
		fmt.Print(fmt.Sprintf("%T: %v", id, id))
		t.FailNow()
	}
}
