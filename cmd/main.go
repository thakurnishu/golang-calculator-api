package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/thakurnishu/golang-calculator-api/middleware"
	"github.com/thakurnishu/golang-calculator-api/route"
)

type Server struct {
	Addr   string
	Router *http.ServeMux
}

func newRouter() *http.ServeMux {
	router := http.NewServeMux()
	return router
}

func (s *Server) newServer() *http.Server {
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", s.Addr),
		Handler: middleware.Logging(s.Router),
	}
	return &server
}

func main() {
	listenAddr, exist := os.LookupEnv("LISTEN_ADDR")
	if !exist {
		slog.Error("Environment variable not defined", "variable", "LISTEN_ADDR")
		os.Exit(1)
	}
	apiUrl, exist := os.LookupEnv("API_URL")
	if !exist {
		slog.Error("Environment variable not defined", "variable", "API_URL")
		os.Exit(1)
	}
	server := Server{
		Addr:   listenAddr,
		Router: newRouter(),
	}

	server.Router.HandleFunc(fmt.Sprintf("GET %s/intro", apiUrl), intro)
	server.Router.HandleFunc(fmt.Sprintf("POST %s/add", apiUrl), route.Addition)
	server.Router.HandleFunc(fmt.Sprintf("POST %s/subtract", apiUrl), route.Subtract)
	server.Router.HandleFunc(fmt.Sprintf("POST %s/multiply", apiUrl), route.Multiply)
	server.Router.HandleFunc(fmt.Sprintf("POST %s/divide", apiUrl), route.Divide)
	server.Router.HandleFunc(fmt.Sprintf("POST %s/sum", apiUrl), route.Sum)

	slog.Info("Starting server", "addr", fmt.Sprintf("%s:%s", apiUrl, listenAddr))
	err := server.newServer().ListenAndServe()
	if err != nil {
		slog.Error("Unable to Start Server", "err", err)
	}
}

func intro(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("This is Calculator API"))
}
