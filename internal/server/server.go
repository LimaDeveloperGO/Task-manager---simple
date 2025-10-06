package server

import (
	"log"
	"net/http"
	"time"

	"task-manager/internal/database"
	"task-manager/internal/handlers"
	"task-manager/internal/repositories/sqlite"
	"task-manager/internal/routes"
	"task-manager/internal/services"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router      chi.Router
	port        string
	db          *database.DB
	taskHandler *handlers.TaskHandler
}

func NewServer(port string) (*Server, error) {
	db, err := database.NewConnection()
	if err != nil {
		return nil, err
	}

	taskRepo := sqlite.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	return &Server{
		router:      chi.NewRouter(),
		port:        port,
		db:          db,
		taskHandler: taskHandler,
	}, nil
}

func (s *Server) SetupRoutes() {
	routes.SetupRoutes(s.router, s.taskHandler)
}

func (s *Server) Start() error {
	s.SetupRoutes()

	server := &http.Server{
		Addr:         ":" + s.port,
		Handler:      s.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("🚀 Servidor rodando na porta %s", s.port)
	return server.ListenAndServe()
}

func (s *Server) Close() error {
	return s.db.Close()
}