package api

import (
	"html/template"
	"net/http"
)

type server struct {
}

func CreateServer() *server {
	return &server{}
}

func (s *server) InitAPI(addr string) error {
	h := http.NewServeMux()

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("../internal/views/base.html"))
		templ.Execute(w, nil)
	})

	server := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	return server.ListenAndServe()
}
