package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type HealthHandler struct{}

type HealthResponse struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

type TaskResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Message:   "Servidor funcionando perfeitamente!",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *HealthHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message":   "Bem-vindo à API Task Manager!",
		"endpoints": []string{
			"GET /api/v1/health - Health check",
			"GET /api/v1/tasks - Listar tasks",
			"POST /api/v1/tasks - Criar task",
		},
		"timestamp": time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handlers para Tasks (implementação básica)
func (h *HealthHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []TaskResponse{
		{
			ID:          1,
			Title:       "Estudar Go",
			Description: "Aprender fundamentos de Go",
			Completed:   false,
			CreatedAt:   time.Now(),
		},
		{
			ID:          2,
			Title:       "Criar API",
			Description: "Desenvolver API REST com Clean Architecture",
			Completed:   true,
			CreatedAt:   time.Now().Add(-24 * time.Hour),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tasks": tasks,
		"total": len(tasks),
	})
}

func (h *HealthHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task criada com sucesso!",
		"status":  "created",
	})
}

func (h *HealthHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	task := TaskResponse{
		ID:          1,
		Title:       "Task " + id,
		Description: "Descrição da task " + id,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *HealthHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task " + id + " atualizada com sucesso!",
		"status":  "updated",
	})
}

func (h *HealthHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task " + id + " deletada com sucesso!",
		"status":  "deleted",
	})
}