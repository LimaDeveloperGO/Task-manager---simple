package routes

import (
	"net/http"

	"task-manager/internal/handlers"
	"task-manager/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router, taskHandler *handlers.TaskHandler) {
	middleware.SetupMiddleware(r)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"ok","message":"API funcionando"}`))
		})

		r.Route("/tasks", func(r chi.Router) {
			r.Get("/", taskHandler.GetTasks)
			r.Post("/", taskHandler.CreateTask)
			r.Get("/{id}", taskHandler.GetTask)
			r.Put("/{id}", taskHandler.UpdateTask)
			r.Delete("/{id}", taskHandler.DeleteTask)
		})
	})
}