package simplehttp

import (
	"io/ioutil"
	"net/http"
)

// GetResponeBody return the resp.Body() string using GET
func GetResponeBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}
