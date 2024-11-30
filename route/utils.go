package route

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

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

func parseResultFloat64(result float64) (float64, error) {
	f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func throwResponseError(w http.ResponseWriter, message string, err error) bool {
	if err != nil {
		http.Error(w, message, http.StatusBadRequest)
		slog.Error(message, "err", err)
		return true
	}
	return false
}
