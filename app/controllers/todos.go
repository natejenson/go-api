package controllers

import (
	"github.com/natejenson/go-api/app/services"
	"github.com/nu7hatch/gouuid"
	"github.com/revel/revel"
)

// The Todos controller
type Todos struct {
	*revel.Controller
}

// GetAll todos
func (c Todos) GetAll() revel.Result {
	t := services.GetTodoRepo().GetAll()
	return c.RenderJSON(t)
}

// Get a todo by id
func (c Todos) Get(id string) revel.Result {
	uuid, _ := parseUUID(id)
	todo, _ := services.GetTodoRepo().Get(uuid)
	return c.RenderJSON(todo)
}

func parseUUID(id string) (uuid.UUID, error) {
	u, err := uuid.ParseHex(id)
	return *u, err
}
