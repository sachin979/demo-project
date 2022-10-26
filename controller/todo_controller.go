package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/sachin979/todo/models"
)

// Get all todos
func FindTodos(c *gin.Context) {
  var todos []models.Todo
  status := c.Query("status")
  if status != "" {
		models.DB.Where("status = ?", status).Find(&todos)
		c.JSON(http.StatusOK, gin.H{"data": todos})
  }  else {
	models.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})
  }
}



func CreateTodo(c *gin.Context) {
	// Validate input
	var input models.CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	// Create todo
	todo := models.Todo{Text: input.Text, Status: input.Status}
	models.DB.Create(&todo)
  
	c.JSON(201, gin.H{"data": todo})
  }

//Find one todo
func FindTodo(c *gin.Context) {  // Get model if exist
	var todo models.Todo

	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
  }



// Update a Todo
func UpdateTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	// Validate input
	var input models.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	models.DB.Model(&todo).Updates(input)
  
	c.JSON(http.StatusOK, gin.H{"data": todo})
  }


  func DeleteTodo(c *gin.Context) {
	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&todo)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }