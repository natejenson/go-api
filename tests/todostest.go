package tests

import (
	"encoding/json"

	"github.com/natejenson/go-api/app/models"
	"github.com/natejenson/go-api/app/services"
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	services.GetTodoRepo().DeleteAll()
}

func (t *AppTest) TestUpPage() {
	t.Get("/")
	t.AssertOk()
	t.Get("/up")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestGetAll() {
	t.Get("/todos")
	t.AssertOk()
	v := make([]models.Todo, 0)
	json.Unmarshal(t.ResponseBody, &v)
	t.Assert(len(v) == 0)

	todo := models.Todo{Title: "test todo"}
	services.GetTodoRepo().Create(todo)

	t.Get("/todos")
	t.AssertOk()
	v = make([]models.Todo, 0)
	json.Unmarshal(t.ResponseBody, &v)
	t.Assert(len(v) == 1)
}
