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
	val, err := config.RedisClient.Get(c, "todos_cache").Result()

	if err == redis.Nil {
		config.DB.Find(&todos)
		data, _ := json.Marshal(todos)
		config.RedisClient.Set(c, "todos_cache", data, 0)
	} else {
		json.Unmarshal([]byte(val), &todos)
	}
	c.JSON(http.StatusOK, todos)
}
