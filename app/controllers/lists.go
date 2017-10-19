package controllers

import (
	"github.com/revel/revel"
	"github.com/natejenson/go-api/app/models"
)

type Lists struct {
	*revel.Controller
}

func (c Lists) Get(id int) revel.Result {
	l := new(List)
	l.id = id
	return c.RenderJSON(l)
}
