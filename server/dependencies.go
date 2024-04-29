package server

import (
	"rate-limiter/controllers"
	"rate-limiter/services"
)

func resolveNotificationController() *controllers.NotificationController {

	controller := &controllers.NotificationController{
		NotificationService: services.NewService(),
	}

	return controller
}
