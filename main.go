package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"os"
)

const defaultPort = 3000

type server struct {
	router *http.ServeMux
	// db *db
}

func (s *server) routes() {
	s.router.HandleFunc("/hc", s.handleHealthCheck())
	s.router.HandleFunc("/index", s.handleIndex())
}

func (s *server) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Still Alive!")
	}
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!"))
	}
}

func (s *server) Run(port int) error {
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), s.router); err != nil {
		return err
	}
	return nil
}

func main() {
	s := server{
		router: http.NewServeMux(),
	}
	s.routes()

	if err := s.Run(defaultPort); err != nil {
		fmt.Println("[ERROR]Fail to setup server")
		os.Exit(1)
	}

	os.Exit(0)
}
