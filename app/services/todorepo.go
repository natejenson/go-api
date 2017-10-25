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

// New creates an initialized todo repository
func New() *InMemoryTodos {
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
		todoRepo = New()
	}
	return todoRepo
}
