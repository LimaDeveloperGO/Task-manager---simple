package services

import (
	"context"
	"fmt"

	"task-manager/internal/models"
	"task-manager/internal/repositories/interfaces"
)

type TaskService struct {
	taskRepo interfaces.TaskRepository
}

func NewTaskService(taskRepo interfaces.TaskRepository) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, req *models.CreateTaskRequest) (*models.TaskResponse, error) {
	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}

	if err := s.taskRepo.Create(ctx, task); err != nil {
		return nil, fmt.Errorf("erro ao criar task: %v", err)
	}

	response := task.ToResponse()
	return &response, nil
}

func (s *TaskService) GetTaskByID(ctx context.Context, id int) (*models.TaskResponse, error) {
	task, err := s.taskRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	response := task.ToResponse()
	return &response, nil
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]*models.TaskResponse, error) {
	tasks, err := s.taskRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*models.TaskResponse
	for _, task := range tasks {
		response := task.ToResponse()
		responses = append(responses, &response)
	}

	return responses, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id int, req *models.UpdateTaskRequest) (*models.TaskResponse, error) {
	// Buscar task existente
	existingTask, err := s.taskRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Atualizar apenas campos fornecidos
	if req.Title != nil {
		existingTask.Title = *req.Title
	}
	if req.Description != nil {
		existingTask.Description = *req.Description
	}
	if req.Completed != nil {
		existingTask.Completed = *req.Completed
	}

	if err := s.taskRepo.Update(ctx, id, existingTask); err != nil {
		return nil, err
	}

	response := existingTask.ToResponse()
	return &response, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.taskRepo.Delete(ctx, id)
}

func (s *TaskService) GetTasksByStatus(ctx context.Context, completed bool) ([]*models.TaskResponse, error) {
	tasks, err := s.taskRepo.GetByCompleted(ctx, completed)
	if err != nil {
		return nil, err
	}

	var responses []*models.TaskResponse
	for _, task := range tasks {
		response := task.ToResponse()
		responses = append(responses, &response)
	}

	return responses, nil
}