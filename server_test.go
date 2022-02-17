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
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status of 200, but got %s", res.Status)
	}
}
