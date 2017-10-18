package controllers

import (
	"github.com/revel/revel"
)

type Up struct {
	*revel.Controller
}

func (c Up) Index() revel.Result {
	return c.Render()
}
