package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo/models"
	"todo/services"

	"github.com/gin-gonic/gin"
)
type Handler struct {
	TodoService services.ITodoService
}

func (h *Handler) ListTodos(c *gin.Context) {
	var filters models.TodoFilterList

	_ = c.ShouldBindQuery(&filters)
	todo, err := h.TodoService.List(filters)

	if err != nil {
		// log.Message("Failed to list Todos")
		c.JSON(500, gin.H{"Error": "Internal Server Error"})
		return
	}
	c.JSON(201, gin.H{"todo": todo})
}

func (h *Handler) CreateTodo(c *gin.Context) {

	var form models.CreateTodoInput
	c.Bind(&form)
	fmt.Println(form)
	todo, err := h.TodoService.Add(&form)
	fmt.Println("todo", todo)
	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry") {
			c.JSON(501, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(201, gin.H{"data": todo})

}

func (h *Handler) GetTodo(c *gin.Context) {
	fmt.Println(c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.TodoService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return
	}
	c.JSON(200, gin.H{"data": todo})
}

func (h *Handler) PutTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form models.PutTodoInput
	c.Bind(&form)
	todo, err := h.TodoService.Put(&form, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": todo})
}

func (h *Handler) PatchTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form models.PatchTodoInput
	c.Bind(&form)
	todoDto, err := h.TodoService.Patch(&form, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": todoDto})
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := h.TodoService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": todo})
}
