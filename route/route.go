package route

import (
	"encoding/json"
	"net/http"
)

func Addition(w http.ResponseWriter, r *http.Request) {

	reqNumbers, err := extractNumbers(r)
	if exist := throwResponseError(w, "Invalid JSON format", err); exist {
		return
	}

	result := reqNumbers.Num1 + reqNumbers.Num2
	f, err := parseResultFloat64(result)
	if exist := throwResponseError(w, "Unable to Parse Value", err); exist {
		return
	}

	resp := ResultResponse{Result: f}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if exist := throwResponseError(w, "Invalid JSON format [Unable to Encode]", err); exist {
		return
	}
}

func Subtract(w http.ResponseWriter, r *http.Request) {
	reqNumbers, err := extractNumbers(r)
	if exist := throwResponseError(w, "Invalid JSON format", err); exist {
		return
	}

	result := reqNumbers.Num1 - reqNumbers.Num2

	f, err := parseResultFloat64(result)
	if exist := throwResponseError(w, "Unable to Parse Value", err); exist {
		return
	}

	resp := ResultResponse{Result: f}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if exist := throwResponseError(w, "Invalid JSON format [Unable to Encode]", err); exist {
		return
	}
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	reqNumbers, err := extractNumbers(r)
	if exist := throwResponseError(w, "Invalid JSON format", err); exist {
		return
	}

	result := reqNumbers.Num1 * reqNumbers.Num2

	f, err := parseResultFloat64(result)
	if exist := throwResponseError(w, "Unable to Parse Value", err); exist {
		return
	}

	resp := ResultResponse{Result: f}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if exist := throwResponseError(w, "Invalid JSON format [Unable to Encode]", err); exist {
		return
	}
}

func Divide(w http.ResponseWriter, r *http.Request) {
	reqNumbers, err := extractNumbers(r)
	if exist := throwResponseError(w, "Invalid JSON format", err); exist {
		return
	}

	if reqNumbers.Num2 == 0 {
		http.Error(w, "num2 cannot be zero", http.StatusOK)
		return
	}
	result := reqNumbers.Num1 / reqNumbers.Num2

	f, err := parseResultFloat64(result)
	if exist := throwResponseError(w, "Unable to Parse Value", err); exist {
		return
	}

	resp := ResultResponse{Result: f}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if exist := throwResponseError(w, "Invalid JSON format [Unable to Encode]", err); exist {
		return
	}
}

func Sum(w http.ResponseWriter, r *http.Request) {
	var numbers SumNumber
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if exist := throwResponseError(w, "Invalid JSON format [Unable to Decode]", err); exist {
		return
	}

	result := 0.0
	for _, num := range numbers.Nums {
		result += num
	}

	f, err := parseResultFloat64(result)
	if exist := throwResponseError(w, "Unable to Parse Value", err); exist {
		return
	}

	resp := ResultResponse{Result: f}
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)
	if exist := throwResponseError(w, "Invalid JSON format [Unable to Encode]", err); exist {
		return
	}
}
