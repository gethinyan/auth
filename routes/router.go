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
		v1.POST("/sendEmail", apiV1.SendEmail)
		v1.POST("/signIn", apiV1.SignIn)
		v1.POST("/signUp", apiV1.SignUp)
		v1.Use(jwt.JWT())
		{
			v1.GET("/users/:id", apiV1.GetUserInfo)
			v1.PUT("/users/:id", apiV1.UpdateUser)
			v1.GET("/memos", apiV1.ListMemo)
			v1.PUT("/memos/:id", apiV1.UpdateMemo)
		}
	}

	return router
}
