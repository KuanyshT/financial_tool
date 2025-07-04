package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KuanyshT/financial-tool/backend/database"
	"github.com/KuanyshT/financial-tool/backend/models"
	"github.com/gorilla/mux"
)

func GetGoals(w http.ResponseWriter, r *http.Request) {

	rows, err := database.DB.Query("SELECT id, title, target_amount, current_amount, created_at FROM goals ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var goals []models.Goal
	for rows.Next() {
		var g models.Goal
		err := rows.Scan(&g.ID, &g.Title, &g.TargetAmount, &g.CurrentAmount, &g.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		goals = append(goals, g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goals)
}

func CreateGoal(w http.ResponseWriter, r *http.Request) {

	var g models.Goal
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	query := `INSERT INTO goals (title, target_amount, current_amount) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := database.DB.QueryRow(query, g.Title, g.TargetAmount, g.CurrentAmount).Scan(&g.ID, &g.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func DeleteGoal(w http.ResponseWriter, r *http.Request) {
	
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid goal ID", 400)
		return
	}

	result, err := database.DB.Exec("DELETE FROM goals WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "Goal not found", 404)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func FullFillGoal(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid goal ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE goals SET current_amount = current_amount + $1 WHERE id = $2", payload.Amount, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func MinusFromGoal(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid goal ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE goals SET current_amount = current_amount - $1 WHERE id = $2", payload.Amount, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
