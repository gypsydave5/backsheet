package backsheet_test

import (
	"backsheet"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServing(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	responseRecorder := httptest.NewRecorder()
	server := backsheet.NewServer(newStubSpreadsheet(""))

	server.ServeHTTP(responseRecorder, request)

	res := responseRecorder.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status of 200, but got %s", res.Status)
	}
}

func TestMissingSheet(t *testing.T) {
	path := "/some-random-path"
	request := httptest.NewRequest(http.MethodGet, path, nil)
	responseRecorder := httptest.NewRecorder()
	server := backsheet.NewServer(newErroringSpreadsheet(errors.New("")))

	server.ServeHTTP(responseRecorder, request)

	res := responseRecorder.Result()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status of 404, but got %s", res.Status)
	}
}

func TestFindingASheet(t *testing.T) {
	want := "I am some JSON"
	spreadsheet := newStubSpreadsheet(want)
	server := backsheet.NewServer(spreadsheet)
	path := "/some-known-path"
	request := httptest.NewRequest(http.MethodGet, path, nil)
	responseRecorder := httptest.NewRecorder()

	server.ServeHTTP(responseRecorder, request)

	res := responseRecorder.Result()
	got := responseBodyToString(t, res)
	if got != want {
		t.Errorf("Expected response body of %s, but got %s", want, got)
	}

	callCount := len(spreadsheet.SheetCalls())
	if callCount != 1 {
		t.Errorf("Expected Sheet to be called once, but was called %s times", callCount)
	}
}

func newStubSpreadsheet(want string) *SpreadsheetMock {
	sheet := &SheetMock{ToJSONFunc: func() string { return want }}
	return &SpreadsheetMock{
		SheetFunc: func(s string) (backsheet.Sheet, error) { return sheet, nil },
	}
}

func responseBodyToString(t *testing.T, res *http.Response) string {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Unexpected error when reading response body: %s", err)
	}
	return string(body)
}

func newErroringSpreadsheet(e error) *SpreadsheetMock {
	return &SpreadsheetMock{SheetFunc: func(s string) (backsheet.Sheet, error) {
		return &SheetMock{}, e
	}}
}
