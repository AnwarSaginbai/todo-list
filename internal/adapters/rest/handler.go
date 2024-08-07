package rest

import (
	"context"
	"github.com/AnwarSaginbai/todo-list/internal/application/domain"
	"github.com/AnwarSaginbai/todo-list/internal/validate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (a *Adapter) PostTaskHandler(c *gin.Context) {
	var input struct {
		Task string `json:"task" validate:"required"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	validationErrors := validate.ValidateTodoStruct(input)
	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	todo := domain.Todo{
		Task:      input.Task,
		Done:      false,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		DeletedAt: nil,
	}

	data, err := a.service.SaveTaskToDB(c, &todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save task"})
		log.Println(data)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (a *Adapter) GetAllHandler(c *gin.Context) {
	data, err := a.service.GetAllTaskFromDB(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		log.Println(data)
		return
	}
	c.JSON(http.StatusOK, data)
}

func (a *Adapter) DeleteTaskHandler(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err = a.service.DeleteTaskFromDB(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "task was successfully deleted"})
}

func (a *Adapter) UpdateTaskHandler(c *gin.Context) {

	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var input struct {
		Task string `json:"task" validate:"required"`
		Done bool   `json:"done"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	validationErrors := validate.ValidateTodoStruct(input)
	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	updatedAt := time.Now()
	todo := domain.Todo{
		Task:      input.Task,
		Done:      input.Done,
		CreatedAt: nil,
		UpdatedAt: &updatedAt,
		DeletedAt: nil,
	}
	if err = a.service.UpdateTaskToDB(context.Background(), id, todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update the task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "successfully updated"})
}
