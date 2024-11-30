package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/thakurnishu/golang-calculator-api/middleware"
)

type Server struct {
    Addr string
    Router *http.ServeMux
}

func newRouter() *http.ServeMux {
    router := http.NewServeMux()
    return router 
}

func (s *Server) newServer() *http.Server {
    server := http.Server{
        Addr: fmt.Sprintf(":%s", s.Addr),
        Handler: middleware.Logging(s.Router),
    }
    return &server 
}

type Numbers struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}
type SumNumber struct {
    Nums []float64 `json:"nums"`
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
        Addr: listenAddr,
        Router: newRouter(),
    }

    server.Router.HandleFunc(fmt.Sprintf("GET %s/intro", apiUrl), intro)
    server.Router.HandleFunc(fmt.Sprintf("POST %s/add", apiUrl), addition)
    server.Router.HandleFunc(fmt.Sprintf("POST %s/subtract", apiUrl), subtract)
    server.Router.HandleFunc(fmt.Sprintf("POST %s/multiply", apiUrl), multiply)
    server.Router.HandleFunc(fmt.Sprintf("POST %s/divide", apiUrl), divide)
    server.Router.HandleFunc(fmt.Sprintf("POST %s/sum", apiUrl), sum)

    slog.Info("Starting server", "addr", fmt.Sprintf("%s:%s", apiUrl,listenAddr))
    err := server.newServer().ListenAndServe()
    if err != nil {
        slog.Error("Unable to Start Server", "err", err)
    }
}

func intro(w http.ResponseWriter,_ *http.Request) {
    w.Write([]byte("This is Calculator API"))
}

func extractNumbers(r *http.Request) (Numbers, error) {
    var nums Numbers
    reqBody := r.Body
    defer r.Body.Close()

    err := json.NewDecoder(reqBody).Decode(&nums)
    if err != nil {
        return Numbers{}, err
    }
    return nums, nil
}

func addition(w http.ResponseWriter,r *http.Request) {

    reqNumbers, err := extractNumbers(r)
    if err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        slog.Error("Invalid JSON format", "err", err)
        return
    }
    result := reqNumbers.Num1 + reqNumbers.Num2
    f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64) 
    response := map[string]float64{"result": f}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func subtract(w http.ResponseWriter,r *http.Request) {
    reqNumbers, err := extractNumbers(r)
    if err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        slog.Error("Invalid JSON format", "err", err)
        return
    }
    result := reqNumbers.Num1 - reqNumbers.Num2

    f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64) 
    response := map[string]float64{"result": f}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func multiply(w http.ResponseWriter,r *http.Request) {
    reqNumbers, err := extractNumbers(r)
    if err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        slog.Error("Invalid JSON format", "err", err)
        return
    }
    result := reqNumbers.Num1 * reqNumbers.Num2

    f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64) 
    response := map[string]float64{"result": f}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func divide(w http.ResponseWriter,r *http.Request) {
    reqNumbers, err := extractNumbers(r)
    if err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        slog.Error("Invalid JSON format", "err", err)
        return
    }
    if reqNumbers.Num2 == 0 {
        fmt.Fprintln(w, "num2 not be zero")
        return
    }
    result := reqNumbers.Num1 / reqNumbers.Num2

    f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64) 
    response := map[string]float64{"result": f}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sum(w http.ResponseWriter,r *http.Request) {
    var numbers SumNumber
    err := json.NewDecoder(r.Body).Decode(&numbers)
    if err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        slog.Error("Invalid JSON format", "err", err)
        return
    }

    result := 0.0
    for _, num := range numbers.Nums {
        result += num
    }

    f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64) 
    response := map[string]float64{"result": f}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
