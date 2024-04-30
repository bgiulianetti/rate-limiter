package server

import (
	"rate-limiter/controllers"
	"rate-limiter/middlewares"

	"github.com/gin-gonic/gin"
)

func mapUrlsToControllers(router *gin.Engine, notificationController *controllers.NotificationController) {

	router.GET("/ping", notificationController.Pong)
	router.POST("notifications/:type/users/:user_id",
		middlewares.AdaptHandler(notificationController.ValidateNotificationType),
		middlewares.AdaptHandler(notificationController.ValidateUserID),
		notificationController.SendNotification)
}
