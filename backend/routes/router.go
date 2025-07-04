package routes

import (
	"github.com/KuanyshT/financial-tool/backend/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/transactions", handlers.GetTransactions).Methods("GET")
	r.HandleFunc("/api/transactions", handlers.CreateTransaction).Methods("POST")
	r.HandleFunc("/api/transactions/{id}", handlers.DeleteTransaction).Methods("DELETE")
	r.HandleFunc("/api/goals", handlers.GetGoals).Methods("GET")
	r.HandleFunc("/api/goals", handlers.CreateGoal).Methods("POST")
	r.HandleFunc("/api/goals/{id}", handlers.DeleteGoal).Methods("DELETE")
	r.HandleFunc("/api/goals/{id}/add", handlers.FullFillGoal).Methods("PATCH")
	r.HandleFunc("/api/goals/{id}/minus", handlers.MinusFromGoal).Methods("PATCH")

	return r
}
