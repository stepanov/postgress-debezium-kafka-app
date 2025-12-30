package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stepanov/postgress-debezium-kafka-app/internal/handlers"
	"github.com/stepanov/postgress-debezium-kafka-app/internal/model"
	mockrepo "github.com/stepanov/postgress-debezium-kafka-app/internal/repository/mock"
)

func TestCreateAndGetUser(t *testing.T) {
	r := chi.NewRouter()
	repo := mockrepo.New()
	h := handlers.NewUsersHandler(repo)
	h.Register(r)

	// create
	payload := map[string]string{"name": "Alice", "email": "alice@example.com"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/users/", bytes.NewReader(b))
	rw := httptest.NewRecorder()
	r.ServeHTTP(rw, req)
	if rw.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d body=%s", rw.Code, rw.Body.String())
	}
	var u model.User
	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		t.Fatalf("decode: %v", err)
	}

	// get
	getReq := httptest.NewRequest(http.MethodGet, "/users/"+u.ID.String(), nil)
	getRw := httptest.NewRecorder()
	r.ServeHTTP(getRw, getReq)
	if getRw.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", getRw.Code)
	}
	var u2 model.User
	if err := json.NewDecoder(getRw.Body).Decode(&u2); err != nil {
		t.Fatalf("decode get: %v", err)
	}
	if u2.Email != u.Email || u2.Name != u.Name {
		t.Fatalf("mismatch")
	}
}
