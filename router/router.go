package router

import (
	"github.com/DualVectorFoil/Zelda/app/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	h := handler.GetHandlerInstance()

	router := gin.Default()
	router.NoRoute(notFound)

	userRoute := router.Group("/user")
	userRoute.POST("/login", h.UserCtrl.Login)
	userRoute.POST("/register", h.UserCtrl.Register)
	userRoute.POST("/verify_code", h.UserCtrl.VerifyCode)
	userRoute.POST("/modify_password", h.UserCtrl.ModifyPassword)

	router.Run(":8080")
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 page not found",
	})
}
