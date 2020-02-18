package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var tipoRepo string

func createTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var newTask Item

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := createRepository()
	repo.CreateItem(newTask)
	c.JSON(http.StatusOK, gin.H{"message:": newTask})
}

func getTasks(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, gin.H{"message": createRepository().GetItems()})
}

func updateTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var task Item

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := createRepository()
	repo.UpdateItem(task)

	c.JSON(http.StatusOK, gin.H{"message": "Tarea actualizada con exito!"})
}

func deleteTask(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var idTask int

	if err := c.ShouldBindQuery(&idTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repo := createRepository()
	repo.DeleteItem(idTask)

	c.JSON(http.StatusOK, gin.H{"message": "Tarea Eliminada con exito!"})

}

func createRepository() TypeRepository {
	if tipoRepo == "mongo" {
		return &MongoRepository{}
	}
	return &MemoryRepository{}
}

func main() {

	tipoRepo = "mongo"
	router := gin.Default()

	api := router.Group("/api")

	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.JSON(http.StatusOK, gin.H{"message": "Bienvenido!"})

	})

	api.POST("/", createTask)
	api.GET("/", getTasks)
	api.PATCH("/:id", updateTask)
	api.DELETE("/:id", deleteTask)

	router.Run(":8080")
}
