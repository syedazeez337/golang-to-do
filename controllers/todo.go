package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/syedazeez337/golang-to-do/config"
	"github.com/syedazeez337/golang-to-do/models"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	val, err := config.RedisClient.Get(config.Ctx, "todos_cache").Result() // Use the global context

	if err == redis.Nil {
		config.DB.Find(&todos)
		data, _ := json.Marshal(todos)
		config.RedisClient.Set(config.Ctx, "todos_cache", data, 0) // Use context here as well
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cache"})
		return
	} else {
		json.Unmarshal([]byte(val), &todos)
	}
	c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
    // Retrieve the ID from the request parameters
    id := c.Param("id")

    // Find the existing TODO item by ID
    var todo models.Todo
    if err := config.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    // Bind the JSON input to the existing TODO struct
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save the updated TODO item to the database
    if err := config.DB.Save(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Todo"})
        return
    }

    // Respond with the updated TODO
    c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
    // Create a new instance of the Todo model
    var todo models.Todo

    // Bind the JSON input to the Todo struct
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save the new Todo to the database
    if err := config.DB.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Todo"})
        return
    }

    // Respond with the created Todo
    c.JSON(http.StatusCreated, todo)
}

func DeleteTodo(c *gin.Context) {
    // Retrieve the ID from the request parameters
    id := c.Param("id")

    // Find the TODO item by ID
    var todo models.Todo
    if err := config.DB.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    // Delete the TODO item from the database
    if err := config.DB.Delete(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Todo"})
        return
    }

    // Respond with a success message
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}