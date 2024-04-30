package server

import (
	"rate-limiter/controllers"
	"rate-limiter/dao"
	"rate-limiter/services"
)

func resolveNotificationController() *controllers.NotificationController {
	controller := &controllers.NotificationController{
		RateLimitService: services.NewRateLimitService(
			dao.NewNotificationContainer(),
			services.NewRulesService(
				dao.NewRulesContainer(),
			),
		),
	}
	return controller
}
