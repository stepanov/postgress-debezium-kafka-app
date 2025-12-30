// Package handlers provides HTTP handlers for the API.
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stepanov/postgress-debezium-kafka-app/internal/model"
	"github.com/stepanov/postgress-debezium-kafka-app/internal/repository"
)

// UsersHandler handles user HTTP endpoints.
type UsersHandler struct {
	Repo repository.UserRepository
}

// NewUsersHandler creates a new UsersHandler.
func NewUsersHandler(r repository.UserRepository) *UsersHandler {
	return &UsersHandler{Repo: r}
}

// Register registers the users routes on the provided router.
func (h *UsersHandler) Register(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.create)
		r.Get("/", h.list)
		r.Get("/{id}", h.get)
		r.Put("/{id}", h.update)
		r.Delete("/{id}", h.delete)
	})
}

func (h *UsersHandler) create(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var payload struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := &model.User{
		ID:        uuid.New(),
		Name:      payload.Name,
		Email:     payload.Email,
		CreatedAt: time.Now().UTC(),
	}
	if err := h.Repo.Create(ctx, u); err != nil {
		h.respondErr(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(u)
}

func (h *UsersHandler) get(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	idStr := chi.URLParam(req, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.respondErr(w, err)
		return
	}
	u, err := h.Repo.GetByID(ctx, id)
	if err != nil {
		h.respondErr(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(u)
}

func (h *UsersHandler) list(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	users, err := h.Repo.List(ctx)
	if err != nil {
		h.respondErr(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}

func (h *UsersHandler) update(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	idStr := chi.URLParam(req, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.respondErr(w, err)
		return
	}
	var payload struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := &model.User{
		ID:    id,
		Name:  payload.Name,
		Email: payload.Email,
	}
	if err := h.Repo.Update(ctx, u); err != nil {
		h.respondErr(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *UsersHandler) delete(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	idStr := chi.URLParam(req, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.respondErr(w, err)
		return
	}
	if err := h.Repo.Delete(ctx, id); err != nil {
		h.respondErr(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *UsersHandler) respondErr(w http.ResponseWriter, err error) {
	hStatus := http.StatusInternalServerError
	// For now all errors are 500; extend with sentinel errors as needed
	hw := map[string]string{"error": err.Error()}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(hStatus)
	_ = json.NewEncoder(w).Encode(hw)
}
