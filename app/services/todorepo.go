package services

import (
	"errors"

	"github.com/natejenson/go-api/app/models"
	"github.com/nu7hatch/gouuid"
)

// TodoRepo is a repository for Todos
type TodoRepo interface {
	Get(id uuid.UUID) (*models.Todo, error)
	GetAll() []*models.Todo
	Create(todo models.Todo) *uuid.UUID
}

// InMemoryTodos represents a concurrent, in-memory repository of Todos
type InMemoryTodos struct {
	// Keep a *concurrent map of guid --> Todo
	todos map[uuid.UUID]*models.Todo
}

// Get a Todo by ID
func (r *InMemoryTodos) Get(id uuid.UUID) (*models.Todo, error) {
	t, ok := r.todos[id]
	if !ok {
		return nil, errors.New("Todo not found")
	}
	return t, nil
}

// GetAll gets all the Todos
func (r *InMemoryTodos) GetAll() []*models.Todo {
	m := make([]*models.Todo, 0, len(r.todos))
	for _, val := range r.todos {
		m = append(m, val)
	}
	return m
}

// Create makes a Todo and returns its UUID
func (r *InMemoryTodos) Create(todo models.Todo) *uuid.UUID {
	u, _ := uuid.NewV4()
	todo.ID = u.String()
	r.todos[*u] = &todo
	return u
}

// New creates an initialized todo repository
func new() *InMemoryTodos {
	a, _ := uuid.NewV4()
	b, _ := uuid.NewV4()
	c, _ := uuid.NewV4()
	return &InMemoryTodos{
		todos: map[uuid.UUID]*models.Todo{
			*a: &models.Todo{ID: a.String(), Title: "Walk the dog"},
			*b: &models.Todo{ID: a.String(), Title: "Go to Mars"},
			*c: &models.Todo{ID: a.String(), Title: "Go outside"},
		},
	}
}

var todoRepo *InMemoryTodos

// GetTodoRepo gets the TodoRepo singleton instance
func GetTodoRepo() (r TodoRepo) {
	if todoRepo == nil {
		todoRepo = new()
	}
	return todoRepo
}
