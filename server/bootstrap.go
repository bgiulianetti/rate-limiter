package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func bootstrap(router *gin.Engine) {
	fmt.Println("Bootstrap - Starting app...")

	application := resolveNotificationController()
	mapUrlsToControllers(router, application)

	fmt.Println("Bootstrap - Application is up")
}
