package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/controllers"
	"github.com/zaahidali/task_manager_api/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.GET("/", middleware.AuthMiddleware(), controllers.GetTasks)
		taskRoutes.GET("/:id", middleware.AuthMiddleware(), controllers.GetTaskByID)
		taskRoutes.POST("/", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.CreateTask)
		taskRoutes.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.UpdateTask)
		taskRoutes.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.DeleteTask)
	}
	userRoutes := router.Group(("/"))
	{
		userRoutes.POST("/register", controllers.Register)
		userRoutes.POST("/login", controllers.Login)
		userRoutes.PUT("/promote", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.Promote)
	}
	return router
}
