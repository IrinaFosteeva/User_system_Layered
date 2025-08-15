package handler

import (
	"encoding/json"
	"github.com/IrinaFosteeva/User_system_layered/internal/custom_errors"
	"github.com/IrinaFosteeva/User_system_layered/internal/interfaces"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	svc interfaces.MainUserService
}

func NewUserHandler(s interfaces.MainUserService) *UserHandler {
	return &UserHandler{svc: s}
}

func (h *UserHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", h.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", h.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/users", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/users/{id:[0-9]+}", h.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{id:[0-9]+}", h.Delete).Methods(http.MethodDelete)
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := h.svc.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, users, http.StatusOK)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ctx := r.Context()
	user, err := h.svc.GetByID(ctx, id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	jsonResponse(w, user, http.StatusOK)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	user, err := h.svc.Create(ctx, in.Name, in.Email)
	if err != nil {
		if err == custom_errors.ErrInvalidInput {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(w, user, http.StatusCreated)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var in struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	user, err := h.svc.Update(ctx, id, in.Name, in.Email)
	if err != nil {
		if err == custom_errors.ErrInvalidInput {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	jsonResponse(w, user, http.StatusOK)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ctx := r.Context()
	err := h.svc.Delete(ctx, id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func jsonResponse(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
