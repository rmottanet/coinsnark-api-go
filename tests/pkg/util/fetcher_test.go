package util_test

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"coinsnark/api/pkg/util"
)

func TestFetchData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Example data"))
	}))
	defer server.Close()

	data, err := util.FetchData(server.URL)
	if err != nil {
		t.Errorf("Unexpected error when fetching data: %v", err)
	}

	expected := "Example data"
	if string(data) != expected {
		t.Errorf("Unexpected result. Expected: %s, Obtained: %s", expected, string(data))
	}

	_, err = util.FetchData("url-invalid")
	if err == nil {
		t.Errorf("An error was expected when fetching data from an invalid URL")
	}
}
