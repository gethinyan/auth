package routes

import (
	apiV1 "github.com/gethinyan/enterprise/api/v1"
	"github.com/gethinyan/enterprise/middleware/jwt"
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
			v1.POST("/memos/syncMemo", apiV1.SyncMemo)
		}
	}

	return router
}
