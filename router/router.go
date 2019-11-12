package router

import (
	"Zelda/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	h := handler.GetHandlerInstance()

	router := gin.Default()
	router.NoRoute(notFound)
	router.POST("/login", h.UserCtrl.Login)
	router.POST("/register", h.UserCtrl.Register)
	router.GET("/verify_code", h.UserCtrl.VerifyCode)
	router.POST("/modify_password", h.UserCtrl.ModifyPassword)

	router.Run(":8080")
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404 page not found",
	})
}
