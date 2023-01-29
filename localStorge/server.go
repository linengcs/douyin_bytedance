package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type viewHandler struct{}

func (vh *viewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("Error with path %s: %v", path, err)
		w.WriteHeader(404)
		w.Write([]byte("404"))
	}

	if strings.HasSuffix(path, ".html") {
		w.Header().Add("Content-Type", "text/html")
	} else if strings.HasSuffix(path, ".mp4") {
		w.Header().Add("Content-Type", "video/mp4")
	} else if strings.HasSuffix(path, ".jpeg") {
		w.Header().Add("Content-Type", "image/jpeg")
	}

	w.Write(data)
}

func main() {
	http.Handle("/", new(viewHandler))
	http.ListenAndServe(":8080", nil)
}
