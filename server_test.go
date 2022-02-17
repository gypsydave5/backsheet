package backsheet

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServing(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	writer := httptest.NewRecorder()
	Server(writer, request)
	res := writer.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status of 200, but got %s", res.Status)
	}
}

func TestMissingSheet(t *testing.T) {
	path := "/some-random-path"
	request := httptest.NewRequest(http.MethodGet, path, nil)
	writer := httptest.NewRecorder()
	Server(writer, request)
	res := writer.Result()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status of 404, but got %s", res.Status)
	}
}
