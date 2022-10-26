package main

import (
"github.com/gin-gonic/gin"
"github.com/sachin979/todo/models"
"github.com/sachin979/todo/controller"
)

func main() {
  r := gin.Default()

  models.ConnectDatabase() 

  r.GET("/todos", controllers.FindTodos)
  r.POST("/todos", controllers.CreateTodo)
  r.GET("/todo/:id", controllers.FindTodo)
  r.PATCH("/todo/:id", controllers.UpdateTodo) 
  r.PUT("/todo/:id", controllers.UpdateTodo)
  r.DELETE("/todo/:id", controllers.DeleteTodo)


  r.Run()
}