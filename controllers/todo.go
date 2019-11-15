package controllers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
	"log"
	"net/http"

	"gin_workshop/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	// var todo models.Todo
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	fmt.Println("todo", todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func AddTodos(c *gin.Context) {

	var todo models.Todo

	// byte to struct
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := models.InsertTodo(&todo)
	if err != nil {
		log.Println("error")
	} else {
		c.JSON(http.StatusOK, map[string]string{"message": "insert todo success"})
	}
}

func GetAllTodo (c *gin.Context) {
	var todos []models.Todo
	models.GetAllTodo(&todos)
	c.JSON(http.StatusOK, todos)
}

func GetTodo (c *gin.Context) {
	var todo models.Todo
	todo.ID, _ = primitive.ObjectIDFromHex(c.Param("id"))
	models.GetTodo(&todo)
	
	if todo.Title != "" {
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(200, map[string]string{ "message: ": "not found "})
		
	}
	
}

func PutTodo (c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	todo.ID, _ = primitive.ObjectIDFromHex(c.Param("id"))
	models.PutTodo(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo (c *gin.Context) {
	var todo models.Todo

	todo.ID, _ = primitive.ObjectIDFromHex(c.Param("id"))
	err := models.DeleteTodo(&todo)
	if err != nil {
		
	}else {
		c.JSON(http.StatusOK, map[string]string{ "message : ": "delete success"})
	}
	
}


