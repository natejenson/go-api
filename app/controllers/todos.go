package controllers

import (
	"strings"

	"github.com/natejenson/go-api/app/models"
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

// Create a todo item
func (c Todos) Create(todo models.Todo) revel.Result {
	todo.Title = strings.TrimSpace(todo.Title)
	if len(todo.Title) == 0 {
		return c.errResponse(400, "A todo must have a title.")
	}
	u := services.GetTodoRepo().Create(todo)
	return c.response(201, u.String())
}

// Edit a todo item
func (c Todos) Edit(id string, todo models.Todo) revel.Result {
	uuid, _ := parseUUID(id)
	t, _ := services.GetTodoRepo().Get(uuid)

	if t == nil {
		return c.errResponse(404, "The todo could not be found.")
	}

	err := services.GetTodoRepo().Overwrite(uuid, todo)
	if err != nil {
		return c.errResponse(400, err.Error())
	}
	return c.response(204, nil)
}

func parseUUID(id string) (uuid.UUID, error) {
	u, err := uuid.ParseHex(id)
	return *u, err
}

func (c Todos) response(status int, o interface{}) revel.Result {
	res := make(map[string]interface{})
	res["data"] = o
	res["error"] = nil
	c.Response.Status = status
	return c.RenderJSON(res)
}

func (c Todos) errResponse(status int, errMsg string) revel.Result {
	res := make(map[string]interface{})
	res["data"] = nil
	res["error"] = errMsg
	c.Response.Status = status
	return c.RenderJSON(res)
}
