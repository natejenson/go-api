package services

import (
	"github.com/natejenson/go-api/app/models"
)

// TodoRepo is a repository for Todos
type TodoRepo struct {
	// Keep a *concurrent map of guid --> Todo
}

// Get a Todo by ID
func (r TodoRepo) Get(id int) models.Todo {
	l := new(models.Todo)
	l.ID = id
	l.Done = false
	l.Title = "Take Bear on a walk"
	return *l
}

// Get all the Todo
func (r TodoRepo) GetAll() []models.Todo {
	t := []models.Todo{r.Get(123), r.Get(456)}
	return t
}
