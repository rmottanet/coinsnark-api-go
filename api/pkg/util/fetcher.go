package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func FetchData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the body of the answer: %v", err)
	}

	return data, nil
}
