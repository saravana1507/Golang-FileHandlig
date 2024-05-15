package router

import (
	"golang-csv/controller"

	"github.com/gin-gonic/gin"
)

func Router(user controller.UserController) *gin.Engine {
	router := gin.Default()
	router.GET("/get", user.FindAll)
	router.GET("/post", user.Save)
	return router
}
