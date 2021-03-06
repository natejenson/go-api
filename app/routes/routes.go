// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tTodos struct {}
var Todos tTodos


func (_ tTodos) GetAll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Todos.GetAll", args).URL
}

func (_ tTodos) Get(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Todos.Get", args).URL
}

func (_ tTodos) Create(
		todo interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "todo", todo)
	return revel.MainRouter.Reverse("Todos.Create", args).URL
}

func (_ tTodos) Edit(
		id string,
		todo interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "todo", todo)
	return revel.MainRouter.Reverse("Todos.Edit", args).URL
}

func (_ tTodos) Delete(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Todos.Delete", args).URL
}


type tUp struct {}
var Up tUp


func (_ tUp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Up.Index", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


