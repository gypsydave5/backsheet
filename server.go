package backsheet

import (
	"net/http"
)

func Server(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		return
	}
	response.WriteHeader(http.StatusNotFound)
}
