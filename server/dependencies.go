package server

import (
	"rate-limiter/controllers"
	"rate-limiter/dao"
	"rate-limiter/services"
)

func resolveNotificationController() *controllers.NotificationController {

	controller := &controllers.NotificationController{
		NotificationService: services.NewNotificationService(
			dao.NewNotificationContainer(),
			dao.NewRulesContainer(),
		),
		RulesService: services.NewRulesService(
			dao.NewRulesContainer(),
		),
	}
	return controller
}
