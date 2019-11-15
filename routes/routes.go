package routes

import (
	"gin_workshop/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{	
		v1.DELETE("todo/:id", controllers.DeleteTodo)
		v1.PUT("todo/:id", controllers.PutTodo)
		v1.GET("todo/:id", controllers.GetTodo)
		v1.GET("todos", controllers.GetAllTodo)
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.AddTodos)
	}

	return r
}
