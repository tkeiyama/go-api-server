package routes

import (
	"api-server/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets routes.
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("user", controllers.GetAllUsers)
		v1.POST("user", controllers.CreateUser)
		v1.GET("user/:id", controllers.GetUserByID)
		v1.PUT("user/:id", controllers.UpdateUser)
		v1.DELETE("user/:id", controllers.DeleteUser)
	}
	return r
}
