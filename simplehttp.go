package simplehttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Identity contains information about the network id
type Identity struct {
	Hostname, ExternalIP string
}

func (i Identity) String() string {
	return fmt.Sprintf("%s (%s)", i.Hostname, i.ExternalIP)
}

const (
	ipbinurl = "http://httpbin.org/ip"
)

// GetResponseBody return the resp.Body() string using GET
func GetResponseBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// GetExternalIP returns the external ip4 or ip6 adress using httpbin.org
func GetExternalIP() (string, error) {
	ipjson, err := GetResponseBody(ipbinurl)
	if err != nil {
		return "127.0.0.1", err
	}
	var ipresp struct {
		Origin string `json:"origin"`
	}
	err = json.Unmarshal(ipjson, &ipresp)
	return ipresp.Origin, err
}

// GetWhoAmI in the form HOSTNAME (EXTERNAL-IP)
func GetWhoAmI() (Identity, error) {
	h, err := os.Hostname()
	if err != nil {
		return Identity{"localhost", "127.0.0.1"}, err
	}
	i, err := GetExternalIP()
	return Identity{h, i}, err
}

// GetJSONMap Unmarshals the Response body to the interface provided
func GetJSONMap(url string) (v map[string]interface{}, err error) {
	bytes, err := GetResponseBody(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &v)
	return
}
