package controllers

import (
	"github.com/natejenson/go-api/app/services"
	"github.com/revel/revel"
)

type Todos struct {
	*revel.Controller
}

func (c Todos) GetAll() revel.Result {
	r := services.TodoRepo{}
	return c.RenderJSON(r.GetAll())
}

func (c Todos) Get(id int) revel.Result {
	r := services.TodoRepo{}
	l := r.Get(id)
	return c.RenderJSON(l)
}
