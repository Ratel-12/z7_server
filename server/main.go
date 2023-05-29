package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{"1", "Coke", 3.52},
	{"2", "Cod", 10.20},
	{"3", "Bread", 2.15},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/pay", ProcessPayment).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST"},
	})

	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var cart []Product

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	totalCost := 0.0
	for _, product := range cart {
		totalCost += product.Price
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment successful!"))
	fmt.Printf("Someone just bought items worth: %.2f in total.\n", totalCost)
}
