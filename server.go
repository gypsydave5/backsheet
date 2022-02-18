package backsheet

import "net/http"

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

func NewServer(spreadsheet Spreadsheet) Server {
	return Server{spreadsheet}
}
