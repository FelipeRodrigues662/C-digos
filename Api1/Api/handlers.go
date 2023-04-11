package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/FelipeRodrigues662/ApiRestGo/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func RegisterHandlers(router *mux.Router, db *gorm.DB) {
	h := &Handler{db}

	router.HandleFunc("/users", h.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", h.GetUserByID).Methods("GET")
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", h.DeleteUser).Methods("DELETE")
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	h.db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := &models.User{}
	result := h.db.First(user, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	h.db.Create(user)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := &models.User{}
	result := h.db.First(user, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found")
		return
	}
	json.NewDecoder(r.Body).Decode(user)
	h.db.Save(user)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := &models.User{}
	result := h.db.First(user, id)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found")
		return
	}
	h.db.Delete(user)
	json.NewEncoder(w).Encode(&models.DeleteResponse{ID: id, Message: "User deleted successfully"})
}
