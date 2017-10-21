package controllers

import (
	"github.com/revel/revel"
	"github.com/natejenson/go-api/app/models"
)

type Lists struct {
	*revel.Controller
}

func (c Lists) Get(id int) revel.Result {
	l := new(models.List)
	l.Id = id
	return c.RenderJSON(l)
}
