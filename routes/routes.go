package routes

import (
	"todo/config"
	controllers "todo/controller"
)

// InitRoutes for the module
func InitRoutes(c config.HandlerConfig) {
	h := controllers.Handler{
		TodoService: c.TodoService,
	}

	const todoIdurl = "/todo/:id"

	g := c.R

	g.GET("/todos", h.ListTodos)
	g.POST("/todos", h.CreateTodo)
	g.GET(todoIdurl, h.GetTodo)
	g.PUT(todoIdurl, h.PutTodo)
	g.PATCH(todoIdurl, h.PatchTodo)
	g.DELETE(todoIdurl, h.DeleteTodo)

}
