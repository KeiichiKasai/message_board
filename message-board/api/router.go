package api

import (
	"github.com/gin-gonic/gin"
	"message-board/api/middleware"
)

func InitRouter() {
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/forget", forget)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.POST("/write_message", WriteMessage)
		UserRouter.POST("/research_message", ResearchMessage)
		UserRouter.POST("/comment", comment)
	}
	r.Run(":8080")
}
