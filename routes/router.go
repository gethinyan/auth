package routes

import (
	apiV1 "e.coding.net/handnote/handnote/api/v1"
	"e.coding.net/handnote/handnote/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/users", apiV1.GetUsers)
		v1.POST("/users", apiV1.CreateUser)
		v1.POST("/sendEmail", apiV1.SendEmail)
		v1.Use(jwt.JWT())
		{
			v1.PUT("/users/:id", apiV1.UpdateUser)
		}
	}

	return router
}
