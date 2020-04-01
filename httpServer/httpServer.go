package main

import (
	"net/http"
	"io/ioutil"
	"path/filepath"
)

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var content []byte
	var err error
	var localPath string

	if req.URL.Path == "/" {
		localPath = "index.html"
		content, err = ioutil.ReadFile(localPath)
	} else {
		localPath = "wwwroot" + req.URL.Path
		content, err = ioutil.ReadFile(localPath)
	}

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}
	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".txt":
		contentType = "text/plain"
	case ".html":
		contentType = "text/html"
	default:
		contentType = "text/plain"
	}
	return contentType
}

func main()  {
	http.Handle("/", new(staticHandler))
	http.ListenAndServe(":80", nil)
	http.ListenAndServe(":8080", nil)
}