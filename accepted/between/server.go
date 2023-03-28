package main

import (
	"io"
	"net/http"
	"strings"
)

type service struct {
	store map[string]string
}

//curl -X POST -d "Kyle XY 33" http://localhost:8080/create
//curl -X GET http://localhost:8080/get

func main() {
	mux := http.NewServeMux()
	srv := service{make(map[string]string)}
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.GetAll)

	http.ListenAndServe("localhost:8080", mux)
}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		splittedContent := strings.Split(string(content), " ")
		s.store[splittedContent[0]] = string(content)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created " + splittedContent[0]))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := ""
		for _, user := range s.store {
			response += user + "\n"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
