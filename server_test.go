package backsheet

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServing(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	responseRecorder := httptest.NewRecorder()
	server := newServer(newStubSpreadsheet(""))
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
	server := newServer(StubMissingSpreadsheet{})
	server.ServeHTTP(responseRecorder, request)
	res := responseRecorder.Result()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status of 404, but got %s", res.Status)
	}
}

func TestFindingASheet(t *testing.T) {
	want := "I am some JSON"
	server := newServer(newStubSpreadsheet(want))
	path := "/some-known-path"
	request := httptest.NewRequest(http.MethodGet, path, nil)
	res := httptest.NewRecorder()
	server.ServeHTTP(res, request)
	body, err := ioutil.ReadAll(res.Result().Body)
	if err != nil {
		t.Errorf("Unexpected error when reading response body: %s", err)
	}
	got := string(body)
	if got != want {
		t.Errorf("Expected response body of %s, but got %s", want, got)
	}
}

func newServer(spreadsheet Spreadsheet) Server {
	return Server{spreadsheet}
}

type Server struct {
	ss Spreadsheet
}

func (s Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	sheet, err := s.ss.Sheet(req.URL.Path)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	res.Write([]byte(sheet.ToJSON()))
}

func newStubSpreadsheet(json string) StubSpreadsheet {
	return StubSpreadsheet{StubSheet{json}}
}

type StubSpreadsheet struct {
	sheet Sheet
}

type StubMissingSpreadsheet struct {
}

func (s StubMissingSpreadsheet) Sheet(_ string) (Sheet, error) {
	return StubSheet{}, errors.New("")
}

func (ss StubSpreadsheet) Sheet(_ string) (Sheet, error) {
	return ss.sheet, nil
}

type StubSheet struct {
	json string
}

func (s StubSheet) ToJSON() string {
	return s.json
}

type Sheet interface {
	ToJSON() string
}

type Spreadsheet interface {
	Sheet(string) (Sheet, error)
}
