package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KuanyshT/financial-tool/backend/database"
	"github.com/KuanyshT/financial-tool/backend/models"

	"github.com/gorilla/mux"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	rows, err := database.DB.Query("SELECT id, category, title, amount, type, created_at FROM transactions ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.ID, &t.Category, &t.Title, &t.Amount, &t.Type, &t.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		transactions = append(transactions, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var t models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO transactions (category, title, amount, type) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := database.DB.QueryRow(query, t.Category, t.Title, t.Amount, t.Type).Scan(&t.ID, &t.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("DELETE FROM transactions WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
