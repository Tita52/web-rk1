package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Result string `json:"result"`
}

func topKekHandler(w http.ResponseWriter, r *http.Request) {
	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		http.Error(w, "Missing 'number' query parameter", http.StatusBadRequest)
		return
	}

	// Преобразуем строку в целое число
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "Invalid 'number' parameter", http.StatusBadRequest)
		return
	}

	// Определяем результат в зависимости от кратности
	var result string
	if number%3 == 0 && number%5 == 0 {
		result = "TopKek"
	} else if number%3 == 0 {
		result = "Top"
	} else if number%5 == 0 {
		result = "Kek"
	} else {
		result = ""
	}

	response := Response{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/topkek", topKekHandler)

	fmt.Println("Server starting at http://127.0.0.1:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
