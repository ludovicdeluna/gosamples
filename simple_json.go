//usr/bin/env go run "$0" "$@" ; exit "$?"
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Profile struct {
	Name    string   `json:"name"`
	Hobbies []string `json:"hobbies"`
}

type JsonServer struct {
	data interface{}
	code int
}

func (j *JsonServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if stop := setHeaders(w, r, j.code); stop {
		return
	}

	// Write body
	if err := json.NewEncoder(w).Encode(&j.data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func setHeaders(w http.ResponseWriter, r *http.Request, code int) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		return true
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	return false
}

func startServer(params string) error {
	log.Printf("Listen on %s", params)
	return http.ListenAndServe(params, nil)
}

func index() *JsonServer {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	return &JsonServer{data: &profile, code: http.StatusOK}
}

func main() {
	http.Handle("/", index())
	log.Fatal(startServer("127.0.0.1:9000"))
}
